package main

import (
	"os"

	"github.com/paulo-henrique-ph/golang-blockchain/commandline"
)

func main() {
	defer os.Exit(0)

	cli := commandline.CommandLine{}
	cli.Run()
}
