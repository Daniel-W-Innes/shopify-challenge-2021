package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Daniel-W-Innes/shopify-add/controller"
	"github.com/Daniel-W-Innes/shopify-add/proto"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := controller.Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterShopifyAddServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
