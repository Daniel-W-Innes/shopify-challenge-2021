package main

import (
	"fmt"
	"github.com/Daniel-W-Innes/shopify-challenge-2021/controller"
	"github.com/Daniel-W-Innes/shopify-challenge-2021/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
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
