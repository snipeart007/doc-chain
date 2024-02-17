package chain

type Miner interface {
	GetID() string
	Mine(Block) Block
}
