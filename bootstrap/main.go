package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/snipeart007/doc-chain/bootstrap/data_service"
	"github.com/snipeart007/doc-chain/bootstrap/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		log.Fatal(err)
	}

	var options []grpc.ServerOption

	grpcServer := grpc.NewServer(options...)
	pb.RegisterDataServiceServer(grpcServer, server.NewDataServiceServer(BLOCKCHAIN_COMPLEXITY))
	log.Printf("Starting gRPC Server on localhost:%d", PORT)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
