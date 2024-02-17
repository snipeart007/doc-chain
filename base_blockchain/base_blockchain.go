package base_blockchain

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/snipeart007/doc-chain/base_blockchain/admin"
	"github.com/snipeart007/doc-chain/base_blockchain/chain"
)

// base_blockchain.Blockchain is the head of the chain.
// All operations on the chain are supposed to be performed on this object.
// Contains the database, the blockchain, the config data and the admin objects.
type Blockchain struct {
	// base_blockchain.Blockchain.db points to the database of the Blockchain object.
	// It can not be altered directly.
	db *chain.DB

	// base_blockchain.Blockchain.chain points to the blockchain of the Blockchain object.
	// It can not be altered directly.
	chain *chain.Chain

	// base_blockchain.Blockchain.Config points to the config data of the Blockchain object.
	// It can be altered directly but should only be altered only through other functions
	// as the same changes need to be done in the database.
	Config *Config
	Admin  *admin.Admin

	// base_blockchain.Blockchain contains the path to the dump directory.
	// Using the absolute path of the dump directory is recommended.
	Directory string
}

// base_blockchain.CreateBlockchain creates a Blockchain instance with the specified dump directory.
// complexity should only be a uint8
func CreateBlockchain(directory string, complexity uint8) (*Blockchain, error) {
	if lastChar := directory[len(directory)-1:]; lastChar != "/" && lastChar != "\\" {
		directory = directory + "/"
	}
	err := os.MkdirAll(directory+"db", 0755)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(directory+"admin", 0755)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(directory+"chain", 0755)
	if err != nil {
		return nil, err
	}

	db, err := chain.NewDB(directory)
	if err != nil {
		return nil, err
	}

	admin.GenerateRSAPair(directory)
	newAdmin, err := admin.ReadAdmin(directory)
	if err != nil {
		return nil, err
	}

	db.SetConfig("complexity", strconv.Itoa(int(complexity)))

	config := &Config{
		Complexity: string(complexity),
	}

	chain := chain.NewChain(db, complexity)

	return &Blockchain{
		db:        db,
		chain:     chain,
		Config:    config,
		Admin:     newAdmin,
		Directory: directory,
	}, nil
}

// base_blockchain.FromDump reads and returns a Blockchain instance from the specified dump directory.
func FromDump(directory string) (*Blockchain, error) {
	if lastChar := directory[len(directory)-1:]; lastChar != "/" && lastChar != "\\" {
		directory = directory + "/"
	}
	db, err := chain.NewDB(directory)
	if err != nil {
		return nil, err
	}

	jsonConfig, err := db.SerializeConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(jsonConfig, config)
	if err != nil {
		return nil, err
	}

	admin, err := admin.ReadAdmin(directory)
	if err != nil {
		return nil, err
	}

	complexity, err := strconv.Atoi(config.Complexity)
	if err != nil {
		return nil, err
	}

	chain, err := chain.ReadChain(directory, db, uint8(complexity))
	if err != nil {
		return nil, err
	}
	return &Blockchain{
		db:        db,
		chain:     chain,
		Config:    config,
		Admin:     admin,
		Directory: directory,
	}, nil
}
