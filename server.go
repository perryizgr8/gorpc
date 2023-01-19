package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"perryizgr8.com/gorpc/procedures"
)

type server struct {
	procedures.PingServer
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxInt64),
		grpc.KeepaliveParams(keepalive.ServerParameters{Timeout: 5 * time.Second}),
	}
	s := grpc.NewServer(opts...)
	procedures.RegisterPingServer(s, &server{})
	fmt.Println("starting server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Ping(ctx context.Context, req *procedures.EchoRequest) (*procedures.EchoResponse, error) {
	fmt.Printf("Got echo request... %v, replying... Pong!\n", req)
	return &procedures.EchoResponse{Msg: "Pong!"}, nil
}
