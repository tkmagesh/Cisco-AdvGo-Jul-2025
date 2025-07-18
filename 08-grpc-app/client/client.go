package main

import (
	"context"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient("localhost:50052", options)
	if err != nil {
		log.Fatalln(err)
	}

	serviceClientProxy := proto.NewAppServiceClient(clientConn)

	// do add operation
	addReq := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	ctx := context.Background()
	addResp, err := serviceClientProxy.Add(ctx, addReq)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Add Result :", addResp.GetResult())
}
