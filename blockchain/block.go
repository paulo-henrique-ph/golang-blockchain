package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

// DeriveHash makes a new hash based
func (block *Block) DeriveHash() {
	info := bytes.Join([][]byte{block.Data, block.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

// Create a block on the chain
// returns the block
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.DeriveHash()
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
