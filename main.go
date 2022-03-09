package main

import (
	"github.com/snipeart007/go-doc-chain/admin"
	"github.com/snipeart007/go-doc-chain/chain"
)

func main() {
	admin.Instance.ReadRSAPair()

	chain.NewDocument(
		admin.Instance.PublicKey,
		chain.BlockData{
			"msg": "Hello World",
		},
	)
}
