package main

import (
	"os"

	"github.com/paulo-henrique-ph/golang-blockchain/blockchain"
	"github.com/paulo-henrique-ph/golang-blockchain/commandline"
)

func main() {
	defer os.Exit(0)

	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := commandline.CommandLine{Blockchain: chain}
	cli.Run()
}
