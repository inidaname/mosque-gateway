package client

import (
	"log"

	pb "github.com/inidaname/mosque/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient(addr string) pb.AuthServiceClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to user-service: %v", err)
	}
	return pb.NewAuthServiceClient(conn)
}
