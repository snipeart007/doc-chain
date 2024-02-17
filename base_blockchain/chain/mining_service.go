package chain

type MiningService interface {
	MineBlock(block Block, complexity uint8, success chan bool) (Block, Miner)
}
