package chain

import (
	"crypto"
	"crypto/rsa"
	"log"
	"time"
)

type Chain struct {
	chain      []Block
	complexity uint8
}

func NewChain(complexity int32) *Chain {
	return &Chain{
		chain: []Block{
			{
				TimeStamp: time.Now(),
				IsMined:   true,
			}}, complexity: uint8(complexity),
	}
}

func (chain *Chain) LastBlockHash() (string, error) {
	lastBlock := chain.chain[len(chain.chain)-1]
	return lastBlock.Hash()
}

func (chain *Chain) AddBlock(document *Document, publicKey *rsa.PublicKey, signature []byte) {
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, document.Hash(), signature)
	if err != nil {
		log.Print(err)
		return
	}

	lastBlockHash, err := chain.LastBlockHash()
	if err != nil {
		log.Print(err)
		return
	}

	newBlock, err := NewBlock(lastBlockHash, document.Data)
	if err != nil {
		log.Print(err)
		return
	}

	if !newBlock.IsMined {
		newBlock.Mine(chain.complexity)
	}

	chain.chain = append(chain.chain, *newBlock)
}
