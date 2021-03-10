package main

import (
	"context"
	pb "github.com/sewasiindonesia/grpc/slack/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewSlackClient(conn)

	response, err := c.SendSlackMessage(context.Background(), &pb.SendSlackMessageRequest{Message: "Testing from client yeay"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Success)
}
