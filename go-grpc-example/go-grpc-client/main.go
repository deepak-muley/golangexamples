// Package main implements a client for Ping service.
package main

import (
	"context"
	"log"
	"time"

	pb "golangexamples/go-grpc-example/pingpong-proto/generated/grpc"

	"google.golang.org/grpc"
)

const (
	address        = "localhost:10001"
	defaultMessage = "ping..."
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendPing(ctx, &pb.PingRequest{Message: defaultMessage})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Pong: %s", r.GetMessage())
}
