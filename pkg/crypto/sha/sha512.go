package sha

import (
	"crypto/sha512"
	"encoding/hex"
)

type SHA512 struct {
	bytes [64]byte
}

func SHA512Hash(data []byte) *SHA512 {
	return &SHA512{
		bytes: sha512.Sum512(data),
	}
}

func (hash *SHA512) ToHex() string {
	return hex.EncodeToString(hash.bytes[:])
}

func (hash *SHA512) Bytes() []byte {
	return hash.bytes[:]
}
