package chain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"
)

type Block struct {
	PrevHash  string    `json:"prevHash"`
	TimeStamp time.Time `json:"timeStamp"`
	Nonce     int32     `json:"nonce"`
	IsMined   bool      `json:"isMined"`
	Data      BlockData `json:"data"`
}

func NewBlock(prevHash string, data BlockData) (*Block, error) {
	randBig, err := rand.Int(rand.Reader, big.NewInt(9999999999))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	nonce := int32(randBig.Int64())
	return &Block{
		PrevHash:  prevHash,
		TimeStamp: time.Now(),
		Nonce:     nonce,
		IsMined:   false,
		Data:      data,
	}, nil
}

func (block *Block) Mine(complexity uint8) int32 {
	var solution int32 = 1
	for !block.IsMined {
		hash := sha256.Sum256([]byte(fmt.Sprintf("%v", block.Nonce+solution)))
		if hex.EncodeToString(hash[:])[:complexity] == strings.Repeat("0", int(complexity)) {
			block.IsMined = true
		}
		solution++
	}
	return solution
}

func (block *Block) Hash() (string, error) {
	jsonString, err := json.Marshal(block)
	if err != nil {
		log.Print(err)
		return "", err
	}

	hash := sha256.Sum256(jsonString)
	return hex.EncodeToString(hash[:]), nil
}
