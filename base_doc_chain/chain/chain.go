package chain

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Chain struct {
	chain      []Block
	complexity uint8
	db         *DB
}

func NewChain(db *DB, complexity uint8) *Chain {
	return &Chain{
		chain: []Block{
			{
				TimeStamp: time.Now(),
				IsMined:   true,
			}},
		complexity: uint8(complexity),
		db:         db,
	}
}

func ReadChain(directory string, db *DB, complexity uint8) (*Chain, error) {
	file, err := os.Open(directory + "chain/chain.gz")
	if err != nil {
		return nil, err
	}

	gziper, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gziper.Close()

	contents, err := ioutil.ReadAll(gziper)
	if err != nil {
		return nil, err
	}

	var chain []Block
	err = json.Unmarshal(contents, &chain)
	if err != nil {
		return nil, err
	}

	return &Chain{
		chain:      chain,
		db:         db,
		complexity: complexity,
	}, nil
}

func (chain *Chain) LastBlock() Block {
	return chain.chain[len(chain.chain)-1]
}

func (chain *Chain) LastBlockHash() (string, error) {
	lastBlock := chain.chain[len(chain.chain)-1]
	return lastBlock.Hash()
}

func (chain *Chain) AddBlock(document *Document, publicKey *rsa.PublicKey, signature []byte) error {
	hash, err := document.Hash()
	if err != nil {
		return err
	}
	
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, signature)
	if err != nil {
		return err
	}

	lastBlockHash, err := chain.LastBlockHash()
	if err != nil {
		return err
	}

	newBlock, err := NewBlock(document.ID, lastBlockHash, document.Data)
	if err != nil {
		return err
	}

	if !newBlock.IsMined {
		newBlock.Mine(chain.complexity)
	}

	chain.chain = append(chain.chain, *newBlock)

	err = chain.db.InsertDocument(hex.EncodeToString(signature), document)
	if err != nil {
		return err
	}
	return nil
}

func (chain *Chain) GenerateDump() ([]byte, error) {
	jsonChain, err := json.Marshal(chain.chain)
	if err != nil {
		return nil, err
	}
	var jsonBuf bytes.Buffer
	err = json.Compact(&jsonBuf, jsonChain)
	if err != nil {
		return nil, err
	}

	var compressed bytes.Buffer
	gziper, _ := gzip.NewWriterLevel(&compressed, gzip.BestCompression)
	gziper.Header.Name = "chain.gz"
	gziper.Header.ModTime = time.Now()
	_, err = gziper.Write(jsonBuf.Bytes())
	if err != nil {
		return nil, err
	}

	err = gziper.Close()
	if err != nil {
		return nil, err
	}

	return compressed.Bytes(), nil
}

func (chain *Chain) Dump(directory string) error {
	dump, err := chain.GenerateDump()
	if err != nil {
		return err
	}

	err = os.WriteFile(directory+"chain/chain.gz", dump, 0644)
	return err
}
