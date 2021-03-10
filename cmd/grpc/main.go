package main

import (
	"fmt"
	grpc2 "github.com/sewasiindonesia/golang-research/grpc"
	"log"
	"net"

	pb "github.com/sewasiindonesia/grpc/slack/proto"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("GRPC Running on port: 9000" )

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc2.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterSlackServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}