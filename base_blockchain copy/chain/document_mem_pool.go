package chain

type DocumentMemPool struct {
	Size      int
	documents chan Document
	chain     *Chain
}

func NewDocumentMemPool(size int, chain *Chain) *DocumentMemPool {
	return &DocumentMemPool{
		Size:      size,
		documents: make(chan Document, size),
		chain:     chain,
	}
}

func (documentMemPool *DocumentMemPool) IsFull() bool {
	return len(documentMemPool.documents) == cap(documentMemPool.documents)
}

func (documentMemPool *DocumentMemPool) Add(document Document) {
	documentMemPool.documents <- document
}

// func CycleDocumentMemPool(documentMemPool *DocumentMemPool) {
// 	for {
// 		log.Print("Cycling through DocumentMemPool")
// 		if documentMemPool.IsFull() {
// 			<-documentMemPool.isReady
// 			documentMemPool.isReady <- false
// 			log.Print(len(documentMemPool.documents))
// 			log.Print(cap(documentMemPool.documents))
// 			log.Print("DocumentMemPool is full!")
// 			close(documentMemPool.documents)
// 			documentMemPool.chain.AddDocuments(ChanToSlice(documentMemPool.documents))
// 			documentMemPool.documents = make(chan Document, 100)
// 			<-documentMemPool.isReady
// 			documentMemPool.isReady <- true
// 		}

// 		time.Sleep(time.Millisecond * 100)
// 	}
// }
