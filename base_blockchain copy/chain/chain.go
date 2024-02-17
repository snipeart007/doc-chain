package chain

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Chain struct {
	chain           []Block
	complexity      uint8
	documentMemPool *DocumentMemPool
	db              *DB
}

func NewChain(db *DB, complexity uint8) *Chain {
	newChain := &Chain{
		chain: []Block{
			{
				ID:        "Genesis Block",
				TimeStamp: time.Now().Unix(),
			},
		},
		complexity: complexity,
		db:         db,
	}

	newChain.documentMemPool = NewDocumentMemPool(100, newChain)

	// go CycleDocumentMemPool(newChain.documentMemPool)

	return newChain
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

	newChain := &Chain{
		chain:      chain,
		complexity: complexity,
		db:         db,
	}

	newChain.documentMemPool = NewDocumentMemPool(100, newChain)

	// go CycleDocumentMemPool(newChain.documentMemPool)

	return newChain, nil
}

func (chain *Chain) LastBlock() Block {
	return chain.chain[len(chain.chain)-1]
}

func (chain *Chain) LastBlockHash() (string, error) {
	lastBlock := chain.chain[len(chain.chain)-1]
	return lastBlock.Hash()
}

func (chain *Chain) addBlockToChain(block Block) {
	success := make(chan bool, 1)
	if !block.IsMined() {
		go func() {
			for {
				go chain.blockchain.MiningService.MineBlock(block, chain.complexity, success)
				if <-success == true {
					if block.IsMined() {
						break
					}
					continue
				}
			}
			chain.chain = append(chain.chain, block)
		}()
	} else {
		chain.chain = append(chain.chain, block)
	}
}

func (chain *Chain) AddDocumentToMemPool(document *Document, publicKey *rsa.PublicKey, signature []byte) error {
	hash, err := document.Hash()
	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, signature)
	if err != nil {
		return err
	}

	chain.documentMemPool.Add(*document)
	return nil
}

func (chain *Chain) AddDocuments(documents []Document) error {
	jsonDocuments, err := json.Marshal(documents)
	if err != nil {
		return err
	}

	prevBlockHash, err := chain.LastBlockHash()
	if err != nil {
		return err
	}

	hash := sha256.Sum256(jsonDocuments)
	id := hash[:]

	block, err := chain.NewBlock(id, prevBlockHash, documents)
	if err != nil {
		return err
	}

	block.GenerateMerkelRoot()

	chain.addBlockToChain(*block)
	chain.db.InsertBlock(block.ID, block)
	for i, v := range block.Documents {
		v.Block = block.ID
		err = chain.db.InsertDocument(v.ID, &v)
		if err != nil {
			return errors.New(string(v.ID) + ": " + err.Error() + "." + block.ID + "[" + strconv.Itoa(i) + "]")
		}
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

	return os.WriteFile(directory+"chain/chain.gz", dump, 0644)
}
