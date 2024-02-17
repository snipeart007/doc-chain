package server

import "github.com/snipeart007/doc-chain/base_blockchain/chain"

type Miner struct {
	ID string `json:"id"`
}

func (miner *Miner) GetID() string {
	return miner.ID
}

func (miner *Miner) Mine(block chain.Block) chain.Block {
	
}
