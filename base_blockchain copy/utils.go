package base_blockchain

import (
	"encoding/json"

	"github.com/snipeart007/doc-chain/base_blockchain/chain"
)

// base_blockchain.Blockchain.InsertDocument inserts a document to the blockchain and reflects the same in the database.
// It takes map[string]interface{} as input.
func (blockchain *Blockchain) InsertDocument(data chain.BlockData) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return chain.NewDocument(blockchain.Admin, jsonData).AddToChain(blockchain.chain)
}

// base_blockchain.Blockchain.InsertJSONDocument inserts a document to the blockchain and reflects the same in the database.
// It takes json data as input.
func (blockchain *Blockchain) InsertJSONDocument(jsonData []byte) (string, error) {
	return chain.NewDocument(blockchain.Admin, jsonData).AddToChain(blockchain.chain)
}

// base_blockchain.Blockchain.RetreiveDocument retreives a document from the database.
func (blockchain *Blockchain) RetreiveDocument(id string) (chain.Document, error) {
	return blockchain.db.RetreiveDocument(id)
}

// base_blockchain.Blockchain.Dump dumps the blockchain to the dump directory.
// It is recommeded to run this function before ending the process to not lose the data in the blockchain.
func (blockchain *Blockchain) Dump() error {
	return blockchain.chain.Dump(blockchain.Directory)
}

// base_blockchain.Blockchain.AllDocuments returns all documents from the database.
func (blockchain *Blockchain) AllDocuments() ([]chain.Document, error) {
	return blockchain.db.GetAllDocuments()
}

// base_blockchain.Blockchain.RetreiveBlock retreives a specific block from the database.
func (blockchain *Blockchain) RetreiveBlock(id string) (chain.Block, error) {
	return blockchain.db.RetreiveBlock(id)
}

// base_blockchain.Blockchain.AllBlocks returns all blocks from the database.
func (blockchain *Blockchain) AllBlocks() ([]chain.Block, error) {
	return blockchain.db.GetAllBlocks()
}
