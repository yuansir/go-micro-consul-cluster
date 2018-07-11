package main

import (
	"go-micro-consul-cluster/data"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"context"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
	"github.com/micro/go-micro"
	"go-micro-consul-cluster/srv/rate/proto"
)

type stay struct {
	HotelID string
	InDate  string
	OutDate string
}

type Rate struct {
	rateTable map[stay]*rate.RatePlan
}

func (s *Rate) GetRates(ctx context.Context, req *rate.Request, rsp *rate.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceId := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %v", traceId)
	}

	for _, hotelID := range req.HotelIds {
		stay := stay{
			HotelID: hotelID,
			InDate:  req.InDate,
			OutDate: req.OutDate,
		}

		if s.rateTable[stay] != nil {
			rsp.RatePlans = append(rsp.RatePlans, s.rateTable[stay])
		}
	}

	return nil
}

func loadRateTable(path string) map[stay]*rate.RatePlan {
	file := data.MustAsset(path)

	rates := []*rate.RatePlan{}
	if err := json.Unmarshal(file, &rates); err != nil {
		log.Fatalf("Failed to load json %v", err)
	}

	rateTable := make(map[stay]*rate.RatePlan)
	for _, ratePlan := range rates {
		stay := stay{
			HotelID: ratePlan.HotelId,
			InDate:  ratePlan.InDate,
			OutDate: ratePlan.OutDate,
		}
		rateTable[stay] = ratePlan
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.rate"),
	)

	service.Init()

	rate.RegisterRateHandler(service.Server(), &Rate{rateTable: loadRateTable("data/rates.json")})

	service.Run()
}
