package main

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	example "github.com/liqiangblogdemos/grpc-sample/gen"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8123")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	example.RegisterTestServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	fmt.Println("Listen at localhost:8123")
	s.Serve(lis)
}

type server struct {
}

func (s *server) GetFiles(context.Context, *example.TestRequest) (*example.TestResponse, error) {
	panic("implement me")
}

func (s *server) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
