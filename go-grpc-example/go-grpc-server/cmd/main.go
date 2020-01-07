// Package main implements a server for Ping service.
package main

import (
	"context"
	"log"
	"net"

	pb "go-grpc-server/api"

	"google.golang.org/grpc"
	"github.com/elastic/go-elasticsearch/v8"
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

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
	  log.Fatalf("Error creating the client: %s", err)
	}
	
	log.Println(elasticsearch.Version)
	res, err := es.Info()
	if err != nil {
	  log.Fatalf("Error getting response: %s", err)
	}
	
	log.Println(res)

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
