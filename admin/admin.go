package admin

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"sync"
)

var readOnce sync.Once
var key *rsa.PrivateKey

type Admin struct {
	Key       *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func NewAdmin() *Admin {
	return &Admin{}
}

func (admin *Admin) ReadRSAPair() {
	readOnce.Do(func() {
		keyPEM, err := ioutil.ReadFile("admin/admin.rsa")
		if err != nil {
			log.Fatal(err)
		}

		pubPEM, err := ioutil.ReadFile("admin/admin.rsa.pub")
		if err != nil {
			log.Fatal(err)
		}

		block, _ := pem.Decode(keyPEM)
		admin.Key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatal(err)
		}

		block, _ = pem.Decode(pubPEM)
		admin.PublicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			log.Fatal(err)
		}

	})
}
