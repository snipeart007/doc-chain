package admin

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)


type Admin struct {
	Key       *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func ReadAdmin(filename string) (*Admin, error) {
	filename = filename + "admin/admin"
	admin := &Admin{}
	keyPEM, err := ioutil.ReadFile(filename + ".rsa")
	if err != nil {
		return nil, err
	}

	pubPEM, err := ioutil.ReadFile(filename + ".rsa.pub")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyPEM)
	admin.Key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	block, _ = pem.Decode(pubPEM)
	admin.PublicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (admin *Admin) SignSHA256(hash []byte) ([]byte, error) {
	signature, err := rsa.SignPKCS1v15(rand.Reader, admin.Key, crypto.SHA256, hash)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
