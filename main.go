package main

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

// DeriveHash makes a new hash based
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Create a block on the chain
// returns the block
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.DeriveHash()
	return block
}

func main() {
}
