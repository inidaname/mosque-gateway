package client

import (
	"log"

	pb "github.com/inidaname/mosque_location/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMosqueClient(addr string) pb.MosqueServiceClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to user-service: %v", err)
	}
	return pb.NewMosqueServiceClient(conn)
}
