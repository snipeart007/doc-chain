package base_doc_chain

import "github.com/snipeart007/doc_chain/base_doc_chain/chain"

// base_doc_chain.BaseDocChain.InsertDocument inserts a document to the blockchain and reflects the same in the database.
func (docChain *BaseDocChain) InsertDocument(data chain.BlockData) error {
	return chain.NewDocument(docChain.Admin, data).AddToChain(docChain.chain)
}

// base_doc_chain.BaseDocChain.RetreiveDocument retreives a document from the database.
func (docChain *BaseDocChain) RetreiveDocument(id string) (chain.Document, error) {
	return docChain.db.RetreiveDocument(id)
}

// base_doc_chain.BaseDocChain.Dump dumps the blockchain to the dump directory.
// It is recommeded to run this function before ending the process to not lose the data in the blockchain.
func (docChain *BaseDocChain) Dump() {
	docChain.chain.Dump(docChain.Directory)
}
