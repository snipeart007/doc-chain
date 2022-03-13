package base_doc_chain

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/snipeart007/doc_chain/base_doc_chain/admin"
	"github.com/snipeart007/doc_chain/base_doc_chain/chain"
)

// base_doc_chain.BaseDocChain is the head of the chain.
// All operations on the chain are supposed to be performed on this object.
// Contains the database, the blockchain, the config data and the admin objects.
type BaseDocChain struct {
	// base_doc_chain.BaseDocChain.db points to the database of the BaseDocChain object.
	// It can not be altered directly.
	db *chain.DB

	// base_doc_chain.BaseDocChain.chain points to the blockchain of the BaseDocChain object.
	// It can not be altered directly.
	chain *chain.Chain

	// base_doc_chain.BaseDocChain.Config points to the config data of the BaseDocChain object.
	// It can be altered directly but should only be altered only through other functions
	// as the same changes need to be done in the database.
	Config *Config
	Admin  *admin.Admin

	// base_doc_chain.BaseDocChain contains the path to the dump directory.
	// Using the absolute path of the dump directory is recommended.
	Directory string
}

// base_doc_chain.Create creates a BaseDocChain instance with the specified dump directory.
// complexity should only be a uint8
func Create(directory string, complexity uint8) (*BaseDocChain, error) {
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

	return &BaseDocChain{
		db:     db,
		chain:  chain,
		Config: config,
		Admin:  newAdmin,
	}, nil
}

// base_doc_chain.FromDump reads and returns a BaseDocChain instance from the specified dump directory.
func FromDump(directory string) (*BaseDocChain, error) {
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
	return &BaseDocChain{
		db:     db,
		chain:  chain,
		Config: config,
		Admin:  admin,
	}, nil
}
