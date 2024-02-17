package main

import (
	"context"
	"log"
	"strconv"
	"time"

	pb "github.com/snipeart007/doc-chain/bootstrap/data_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(SERVER_ADDR, opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewDataServiceClient(conn)

	start := time.Now()

	for i := 1; i <= 101; i++ {
		client.InsertDocument(
			context.Background(),
			&pb.CreateDocumentArgs{
				Data: []byte("{\"count\": " + strconv.Itoa(i) + "}"),
			},
		)
	}

	log.Print("101 document insertions took " + time.Since(start).String())
}
