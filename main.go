package main

import (
	"fmt"

	"github.com/paulo-henrique-ph/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("1")
	chain.AddBlock("2")
	chain.AddBlock("3")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
	}
}
