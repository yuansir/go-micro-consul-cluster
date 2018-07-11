package main

import (
	"github.com/hailocab/go-geoindex"
	"context"
	"go-micro-consul-cluster/srv/geo/proto"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
	"go-micro-consul-cluster/data"
	"encoding/json"
	"log"
	"github.com/micro/go-micro"
)

const (
	maxSearchRadius  = 10
	maxSearchResults = 5
)

type point struct {
	Pid  string  `json:"hotelId"`
	Plat float64 `json:"lat"`
	Plon float64 `json:"lon"`
}

func (p *point) Lat() float64 { return p.Plat }
func (p *point) Lon() float64 { return p.Plon }
func (p *point) Id() string   { return p.Pid }

type Geo struct {
	index *geoindex.ClusteringIndex
}

func (s *Geo) NearBy(ctx context.Context, req *geo.Request, rsp *geo.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	center := &geoindex.GeoPoint{
		Pid:  "",
		Plat: float64(req.Lat),
		Plon: float64(req.Lon),
	}

	points := s.index.KNearest(center, maxSearchResults, geoindex.Km(maxSearchRadius), func(p geoindex.Point) bool {
		return true
	})

	for _, p := range points {
		rsp.HotelIds = append(rsp.HotelIds, p.Id())
	}
	return nil
}

func newGeoIndex(path string) *geoindex.ClusteringIndex {
	file := data.MustAsset(path)

	var points []*point
	if err := json.Unmarshal(file, &points); err != nil {
		log.Fatalf("Failed to load hotels: %v", err)
	}

	index := geoindex.NewClusteringIndex()
	for _, point := range points {
		index.Add(point)
	}
	return index
}


func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.geo"),
	)

	service.Init()

	geo.RegisterGeoHandler(service.Server(), &Geo{index: newGeoIndex("data/locations.json")})

	service.Run()
}
