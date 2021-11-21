package main

import (
	"github.com/tyomhk2015/gocoin/blockchain"
)

func main() {
	myChain := blockchain.GetBlockchain()
	myChain.AddBlock("Cover")
	myChain.AddBlock("Hololive")
	myChain.AddBlock("0-Gen-Members")
	myChain.ShowAllBlocks()
}
