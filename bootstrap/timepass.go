package main

// import (
// 	"log"
// 	"strconv"

// 	"github.com/snipeart007/doc-chain/base_blockchain"
// 	"github.com/snipeart007/doc-chain/base_blockchain/chain"
// )

// func main() {
// 	blockchain, err := base_blockchain.Create("./test", 4)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i := 1; i <= 101; i++ {
// 		id, _ := blockchain.InsertDocument(chain.BlockData{
// 			"count": strconv.Itoa(i),
// 		})
// 		log.Print(id)
// 	}

// 	block, err := blockchain.AllBlocks()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("%+v", block[0].MerkelRoot)

// 	err = blockchain.Dump()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	blockchain, err := base_blockchain.FromDump("./test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	document, err := blockchain.RetreiveDocument("document#7125253e719a23ddf0852d9832499714efd10d8bdb59ab4e77f9fee533d2311c")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("Data: %+v", document.Data)

// 	jsonData, err := ioutil.ReadFile("glossary.json")
// 	blockchain.InsertDocument()
// 	if err = blockchain.Dump(); err != nil {
// 		log.Fatal(err)
// 	}
// }
