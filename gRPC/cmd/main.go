package main

import (
	"context"
	"log"
	"net"

	app "github.com/meles-z/algorithm/gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	add := ":9292"
	list, err := net.Listen("tcp", add)
	if err != nil {
		log.Fatalf("error: cant listen :%s", err)
	}
	srv := grpc.NewServer()
	var u Rides
	reflection.Register(srv)
	app.RegisterRidesServer(srv, &u)
	log.Printf("Info: server ready on: %s", add)
	if err := srv.Serve(list); err != nil {
		log.Fatalf("Error to serve: %s", err)
	}
}

type Rides struct {
	app.UnimplementedRidesServer
}

func (r Rides) Start(ctx context.Context, req *app.StartRequest) (*app.StartResponse, error) {
	resp := app.StartResponse{
		Id: req.Id,
	}
	return &resp, nil
}
