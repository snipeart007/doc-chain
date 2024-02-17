package chain

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
)

type DB struct {
	DB       *bolt.DB
	FileName string
}

func NewDB(filename string) (*DB, error) {
	filename = filename + "db/bolt.db"
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("documents"))
	if err != nil {
		return nil, err
	}

	_, err = tx.CreateBucketIfNotExists([]byte("blocks"))
	if err != nil {
		return nil, err
	}

	_, err = tx.CreateBucketIfNotExists([]byte("config"))
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &DB{
		DB:       db,
		FileName: filename,
	}, nil
}

func (db *DB) SetConfig(key string, value string) error {
	tx, err := db.DB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.Bucket([]byte("config")).Put([]byte(key), []byte(value))
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (db *DB) SerializeConfig() ([]byte, error) {
	config := make(map[string]string)
	tx, err := db.DB.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	tx.Bucket([]byte("config")).ForEach(func(k, v []byte) error {
		config[string(k)] = string(v)
		return nil
	})

	return json.Marshal(config)
}

func (db *DB) InsertDocument(id string, document *Document) error {
	tx, err := db.DB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	serialized, err := document.Serialize()
	if err != nil {
		return err
	}

	err = tx.Bucket([]byte("documents")).Put([]byte(id), serialized)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *DB) RetreiveDocument(id string) (Document, error) {
	var jsonValue bytes.Buffer
	db.DB.View(func(tx *bolt.Tx) error {
		jsonValue.Write(tx.Bucket([]byte("documents")).Get([]byte(id)))
		return nil
	})

	if jsonValue.Bytes() == nil {
		return Document{}, ErrDocumentDoesNotExist
	}

	document := new(Document)

	if err := json.NewDecoder(&jsonValue).Decode(document); err != nil {
		return Document{}, err
	}
	return *document, nil
}

func (db *DB) GetAllDocuments() ([]Document, error) {
	tx, err := db.DB.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var jsonDocuments [][]byte

	tx.Bucket([]byte("documents")).ForEach(func(k, v []byte) error {
		jsonDocuments = append(jsonDocuments, v)
		return nil
	})

	for _, v := range jsonDocuments {
		log.Print("document " + string(v))
	}

	var documents []Document
	for _, document := range jsonDocuments {
		doc := &Document{}
		err := json.Unmarshal(document, doc)
		if err != nil {
			return nil, err
		}
		documents = append(documents, *doc)
	}

	return documents, nil
}

func (db *DB) InsertBlock(id string, block *Block) error {
	tx, err := db.DB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	serialized, err := block.Serialize()
	if err != nil {
		return err
	}

	err = tx.Bucket([]byte("blocks")).Put([]byte(id), serialized)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *DB) RetreiveBlock(id string) (Block, error) {
	var jsonValue bytes.Buffer
	db.DB.View(func(tx *bolt.Tx) error {
		jsonValue.Write(tx.Bucket([]byte("blocks")).Get([]byte(id)))
		return nil
	})

	block := new(Block)

	if err := json.NewDecoder(&jsonValue).Decode(block); err != nil {
		return Block{}, err
	}
	return *block, nil
}

func (db *DB) GetAllBlocks() ([]Block, error) {
	tx, err := db.DB.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var jsonBlocks [][]byte

	tx.Bucket([]byte("blocks")).ForEach(func(k, v []byte) error {
		jsonBlocks = append(jsonBlocks, v)
		return nil
	})

	var blocks []Block
	for _, jsonBlock := range jsonBlocks {
		block := &Block{}
		err := json.Unmarshal(jsonBlock, block)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, *block)
	}

	return blocks, nil
}
