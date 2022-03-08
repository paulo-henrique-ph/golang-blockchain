package main

import (
	"fmt"
	"strconv"

	"github.com/paulo-henrique-ph/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("1")
	chain.AddBlock("2")
	chain.AddBlock("3")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	}
}
