package admin

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)


func GenerateRSAPair(directory string) error {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	public := &key.PublicKey
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
			Bytes: x509.MarshalPKCS1PublicKey(public),
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile(directory+"admin/admin.rsa", keyPEM, 0700); err != nil {
		return err
	}

	// Write public key to file.
	if err := ioutil.WriteFile(directory+"admin/admin.rsa.pub", pubPEM, 0755); err != nil {
		return err
	}

	return nil
}
