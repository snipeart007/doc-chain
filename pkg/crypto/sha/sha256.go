package sha

import (
	"crypto/sha256"
	"encoding/hex"
)

type SHA256 struct {
	bytes [32]byte
}

func SHA256Hash(data []byte) *SHA256 {
	return &SHA256{
		bytes: sha256.Sum256(data),
	}
}

func (hash *SHA256) ToHex() string {
	return hex.EncodeToString(hash.bytes[:])
}

func (hash *SHA256) Bytes() []byte {
	return hash.bytes[:]
}
