package admin

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"sync"
)

var generatationOnce sync.Once

func GenerateRSAPair() {
	generatationOnce.Do(func() {
		var err error
		key, err = rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			log.Fatal(err)
		}
	})
	public := key.Public()
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(public.(*rsa.PublicKey)),
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile("admin/admin.rsa", keyPEM, 0700); err != nil {
		panic(err)
	}

	// Write public key to file.
	if err := ioutil.WriteFile("admin/admin.rsa.pub", pubPEM, 0755); err != nil {
		panic(err)
	}
}
