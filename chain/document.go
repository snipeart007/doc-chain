package chain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"log"
)

type Document struct {
	PublicKey *rsa.PublicKey `json:"publicKey"`
	Data      BlockData      `json:"data"`
}

func NewDocument(publicKey *rsa.PublicKey, data BlockData) *Document {
	return &Document{
		PublicKey: publicKey,
		Data:      data,
	}
}

func (document *Document) Hash() []byte {
	jsonBytes, err := json.Marshal(document)
	if err != nil {
		log.Print(err)
		return nil
	}

	hash := sha256.Sum256(jsonBytes)
	return hash[:]
}

func (document *Document) AddToChain(chain *Chain, key *rsa.PrivateKey) {
	hash := document.Hash()
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash[:])
	if err != nil {
		log.Print(err)
		return
	}

	chain.AddBlock(document, document.PublicKey, signature)
}
