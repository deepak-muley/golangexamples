// Package main implements a server for Ping service.
package main

import (
	"context"
	"log"
	"net"

	pb "golangexamples/go-grpc-example/pingpong-proto/generated/grpc"

	"google.golang.org/grpc"
)

const (
	port = ":10001"
)

// server is used to implement pingpong.PingServer.
type server struct {
	pb.UnimplementedPingServer
}

// SendPing implements pingpong.PingServer
func (s *server) SendPing(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.PingReply{Message: "Pong " + in.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
