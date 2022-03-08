package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nounce       int
}

// Create a block on the chain
// returns the block
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	pow := NewProof(block)
	nounce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nounce = nounce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
