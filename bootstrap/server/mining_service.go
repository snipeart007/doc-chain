package server

import "github.com/snipeart007/doc-chain/base_blockchain/chain"

type MiningService struct {}

func (miningService *MiningService) MineBlock(block chain.Block, complexity uint8, success chan bool) (chain.Block, chain.Miner) {
	miner := &Miner{
		ID: "mining",
	}

	success <- true
	return block, miner
}
