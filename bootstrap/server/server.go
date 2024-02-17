package server

import (
	"log"

	"github.com/snipeart007/doc-chain/base_blockchain"
	pb "github.com/snipeart007/doc-chain/bootstrap/data_service"
)

type DataServiceServer struct {
	pb.UnimplementedDataServiceServer
	blockchain *base_blockchain.Blockchain
}

func NewDataServiceServer(complexity uint8) *DataServiceServer {
	blockchain, err := base_blockchain.CreateBlockchain("./blockchain", complexity)
	if err != nil {
		log.Panic(err)
	}

	log.Print("NewDataServiceServer: Created new blockchain")

	// for i := 1; i <= 101; i++ {
	// 	id, _ := blockchain.InsertDocument(chain.BlockData{
	// 		"count": strconv.Itoa(i),
	// 	})
	// 	log.Print(id)
	// }

	// err = blockchain.Dump()
	// if err != nil {
	// 	log.Panic(err)
	// }

	return &DataServiceServer{
		blockchain: blockchain,
	}
}

func DataServiceServerFromDump(complexity uint8) *DataServiceServer {
	blockchain, err := base_blockchain.FromDump("./blockchain")
	if err != nil {
		log.Panic(err)
	}

	log.Print("DataServiceServerFromDump: Blockchain has been read from dump")

	return &DataServiceServer{
		blockchain: blockchain,
	}
}
