package chain

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"

	"github.com/snipeart007/doc_chain/base_doc_chain/admin"
)

type Document struct {
	ID        []byte         `json:"-"`
	PublicKey *rsa.PublicKey `json:"publicKey"`
	Data      BlockData      `json:"data"`
	admin     *admin.Admin   `json:"-"`
}

func NewDocument(admin *admin.Admin, data BlockData) *Document {
	return &Document{
		PublicKey: admin.PublicKey,
		Data:      data,
		admin:     admin,
	}
}

func (document *Document) Hash() ([]byte, error) {
	jsonBytes, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(jsonBytes)
	document.ID = hash[:]
	return document.ID, nil
}

func (document *Document) Serialize() ([]byte, error) {
	return json.Marshal(document)
}

func (document *Document) AddToChain(chain *Chain) error {
	hash, err := document.Hash()
	if err != nil {
		return err
	}
	signature, err := document.admin.SignSHA256(hash)
	if err != nil {
		return err
	}

	return chain.AddBlock(document, document.PublicKey, signature)
}
