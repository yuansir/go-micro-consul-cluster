package main

import (
	"go-micro-consul-cluster/srv/profile/proto"
	"go-micro-consul-cluster/srv/rate/proto"
	"github.com/micro/go-micro/client"
	"context"
	"go-micro-consul-cluster/api/hotel/proto"
	"golang.org/x/net/trace"
	"github.com/micro/go-micro/metadata"
	"github.com/nu7hatch/gouuid"
	"github.com/kataras/iris/core/errors"
	"strings"
	"encoding/base64"
	merr "github.com/micro/go-micro/errors"
	"go-micro-consul-cluster/srv/auth/proto"
	"go-micro-consul-cluster/srv/geo/proto"
	"net/http"
	"github.com/micro/go-micro"
)

const (
	BASIC_SCHEMA  string = "Basic "
	BEARER_SCHEMA string = "Bearer "
)

type profileResults struct {
	hotels []*profile.Hotel
	err    error
}

type rateResults struct {
	ratePlans []*rate.RatePlan
	err       error
}

type Hotel struct {
	Client client.Client
}

func (s *Hotel) Rates(ctx context.Context, req *hotel.Reqeust, rsp *hotel.Response) error {
	tr := trace.New("api.v1", "Hotel.Rates")
	defer tr.Finish()

	ctx = trace.NewContext(ctx, tr)

	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}

	if traceID, err := uuid.NewV4(); err == nil {
		tmd := metadata.Metadata{}
		for k, v := range md {
			tmd[k] = v
		}

		tmd["traceID"] = traceID.String()
		tmd["fromName"] = "api.v1"
		ctx = metadata.NewContext(ctx, tmd)
	}

	token, err := getToken(md)
	if err != nil {
		return merr.Forbidden("api.hotel.rates", err.Error())
	}

	authClient := auth.NewAuthService("go.micro.srv.auth", s.Client)
	if _, err = authClient.VerifyToken(ctx, &auth.Request{AuthToken: token}); err != nil {
		return merr.Unauthorized("api.hotel.rates", "Unauthorized")
	}

	inDate, outDate := req.InDate, req.OutDate
	if inDate == "" || outDate == "" {
		return merr.BadRequest("api.hotel.rates", "Please specify inDate/outDate params")
	}

	// finds nearby hotels
	// TODO(hw): use lat/lon from request params
	geoClient := geo.NewGeoService("go.micro.srv.geo", s.Client)
	nearby, err := geoClient.NearBy(ctx, &geo.Request{
		Lat: 51.502973,
		Lon: -0.114723,
	})

	if err != nil {
		return merr.InternalServerError("api.hotel.rates...", err.Error())
	}

	profileCh := getHotelProfiles(s.Client, ctx, nearby.HotelIds)
	rateCh := getRatePlans(s.Client, ctx, nearby.HotelIds, inDate, outDate)

	profileReply := <-profileCh
	if err := profileReply.err; err != nil {
		return merr.InternalServerError("api.hotel.rates!!!", err.Error())
	}

	rateReply := <-rateCh
	if err := rateReply.err; err != nil {
		return merr.InternalServerError("api.hotel.rates????", err.Error())
	}

	rsp.Hotels = profileReply.hotels
	rsp.RatePlans = rateReply.ratePlans
	return nil
}

func getToken(md metadata.Metadata) (string, error) {
	authHeader := md["Authorization"]
	if authHeader == "" {
		return "", errors.New("Authorization header required")
	}

	// Confirm the request is sending Basic Authentication credentials.
	if !strings.HasPrefix(authHeader, BASIC_SCHEMA) && !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		return "", errors.New("Authorization requires Basic/Bearer scheme")
	}

	// Get the token from the request header
	// The first six characters are skipped - e.g. "Basic ".
	if strings.HasPrefix(authHeader, BASIC_SCHEMA) {
		str, err := base64.StdEncoding.DecodeString(authHeader[len(BASIC_SCHEMA):])
		if err != nil {
			return "", errors.New("Base64 encoding issue")
		}
		creds := strings.Split(string(str), ":")
		return creds[0], nil
	}

	return authHeader[len(BEARER_SCHEMA):], nil
}

func getRatePlans(c client.Client, ctx context.Context, hotelIDs []string, inDate string, outDate string) chan rateResults {
	rateClient := rate.NewRateService("go.micro.srv.rate", c)
	ch := make(chan rateResults, 1)

	go func() {
		res, err := rateClient.GetRates(ctx, &rate.Request{
			HotelIds: hotelIDs,
			InDate:   inDate,
			OutDate:  outDate,
		})
		ch <- rateResults{res.RatePlans, err}
	}()

	return ch
}

func getHotelProfiles(c client.Client, ctx context.Context, hotelIDs []string) chan profileResults {
	profileClient := profile.NewProfileService("go.micro.srv.profile", c)
	ch := make(chan profileResults, 1)

	go func() {
		res, err := profileClient.GetProfiles(ctx, &profile.Request{
			HotelIds: hotelIDs,
			Locale:   "en",
		})
		ch <- profileResults{res.Hotels, err}
	}()

	return ch
}

func main() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	service := micro.NewService(
		micro.Name("go.micro.api.hotel"),
	)

	service.Init()
	hotel.RegisterHotelHandler(service.Server(), &Hotel{service.Client()})
	service.Run()
}