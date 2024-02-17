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
	ID         string     `json:"id"`
	PrevHash   string     `json:"prevHash"`
	TimeStamp  int64      `json:"timeStamp"`
	Nonce      int32      `json:"nonce"`
	Proof      string     `json:"proof"`
	Miner      string     `json:"miner"`
	Documents  []Document `json:"documents"`
	MerkelRoot string     `json:"merkelRoot"`
	chain      *Chain
}

func (chain *Chain) NewBlock(id []byte, prevHash string, documents []Document) (*Block, error) {
	randBig, err := rand.Int(rand.Reader, big.NewInt(9999999999))
	if err != nil {
		return nil, err
	}
	nonce := int32(randBig.Int64())
	return &Block{
		ID:        "block#" + hex.EncodeToString(id),
		PrevHash:  prevHash,
		TimeStamp: time.Now().Unix(),
		Nonce:     nonce,
		Documents: documents,
		chain:     chain,
	}, nil
}

func (block *Block) Serialize() ([]byte, error) {
	return json.Marshal(block)
}

func (block *Block) IsMined() bool {
	if block.Proof == "" {
		return false
	} else if block.Proof[:block.chain.complexity] == strings.Repeat("0", int(block.chain.complexity)) {
		return true
	}
	return false
}

func (block *Block) Mine() int32 {
	start := time.Now()
	var solution int32 = 0
	for !block.IsMined() {
		hashBytes := sha256.Sum256([]byte(fmt.Sprintf("%v", fmt.Sprintf("%s%d", block.ID, block.Nonce+solution))))
		hash := hex.EncodeToString(hashBytes[:])
		if block.IsMined() {
			block.Proof = hash
		}
		solution++
	}
	log.Print("Mined in " + time.Since(start).String())
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

func createMerkelRoot(pairs []string) string {
	if len(pairs) == 1 {
		return pairs[0]
	} else if len(pairs)%2 == 1 {
		pairs = append(pairs, pairs[len(pairs)-1])
	}

	branches := make([]string, 0)
	for i := 0; i < len(pairs); i += 2 {
		pair := sha256.Sum256([]byte(pairs[i] + pairs[i+1]))
		branches = append(branches, hex.EncodeToString(pair[:]))
	}

	return createMerkelRoot(branches)
}

func (block *Block) GenerateMerkelRoot() {
	documents := make([]string, 0)
	for _, document := range block.Documents {
		documents = append(documents, document.ID)
	}

	block.MerkelRoot = createMerkelRoot(documents)
}
