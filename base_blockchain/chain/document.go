package chain

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	"github.com/snipeart007/doc-chain/base_blockchain/admin"
)

type Document struct {
	ID        string         `json:"id"`
	PublicKey *rsa.PublicKey `json:"publicKey"`
	Data      []byte         `json:"data"`
	TimeStamp int64          `json:"timestamp"`
	Block     string         `json:"block"`
	admin     *admin.Admin   `json:"-"`
}

func NewDocument(admin *admin.Admin, data []byte) *Document {
	return &Document{
		PublicKey: admin.PublicKey,
		Data:      data,
		TimeStamp: time.Now().Unix(),
		admin:     admin,
	}
}

func (document *Document) Serialize() ([]byte, error) {
	return json.Marshal(document)
}

func (document *Document) Hash() ([]byte, error) {
	tempDocument := document
	document.ID = ""

	jsonBytes, err := json.Marshal(tempDocument)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(jsonBytes)
	document.ID = "document#" + hex.EncodeToString(hash[:])
	return hash[:], nil
}

func (document *Document) AddToChain(chain *Chain) (string, error) {
	hash, err := document.Hash()
	if err != nil {
		return "", err
	}

	signature, err := document.admin.SignSHA256(hash)
	if err != nil {
		return "", err
	}

	if chain.documentMemPool.IsFull() {
		log.Print("DocumentMemPool is full!")
		close(chain.documentMemPool.documents)
		chain.documentMemPool.chain.AddDocuments(ChanToSlice(chain.documentMemPool.documents))
		chain.documentMemPool.documents = make(chan Document, 100)
	}

	err = chain.AddDocumentToMemPool(document, document.PublicKey, signature)
	if err != nil {
		return "", err
	}
	return document.ID, nil
}
