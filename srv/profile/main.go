package main

import (
	"go-micro-consul-cluster/srv/profile/proto"
	"context"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
	"go-micro-consul-cluster/data"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/micro/go-micro"
)


type Profile struct {
	hotels map[string]*profile.Hotel
}

func (s *Profile) GetProfiles(ctx context.Context, req *profile.Request, rsp *profile.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	for _, i := range req.HotelIds {
		rsp.Hotels = append(rsp.Hotels, s.hotels[i])
	}
	return nil
}

func loadProfiles(path string) map[string]*profile.Hotel {
	file := data.MustAsset(path)

	hotels := []*profile.Hotel{}
	if err := json.Unmarshal(file, &hotels); err != nil {
		log.Fatalf("Failed to load json: %v", err)
	}

	profiles := make(map[string]*profile.Hotel)
	for _, hotel := range hotels {
		profiles[hotel.Id] = hotel
	}
	return profiles
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.profile"),
	)

	service.Init()

	profile.RegisterProfileHandler(service.Server(), &Profile{
		hotels: loadProfiles("data/profiles.json"),
	})

	service.Run()
}