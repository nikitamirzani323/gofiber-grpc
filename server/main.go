package main

import (
	"context"
	"net"

	"github.com/nikitamirzani323/gofiber-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedAddServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}

func (s *server) Add(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b, c := request.GetA(), request.GetB(), request.GetC()

	result := a + b + c

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b, c := request.GetA(), request.GetB(), request.GetC()

	result := a * b * c

	return &proto.Response{Result: result}, nil
}
