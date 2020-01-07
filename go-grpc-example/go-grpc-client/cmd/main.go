// Package main implements a client for Ping service.
package main

import (
	"fmt"
	"context"
	"log"
	"time"
	"os"

	pb "go-grpc-client/api"

	"google.golang.org/grpc"
)

const (
	defaultMessage = "ping..."
)

func main() {
	address := fmt.Sprintf("%s:%s", os.Getenv("go-grpc-server"), os.Getenv("go-grpc-server-port"))
	log.Printf("address:%s", address)
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
