package main

import "github.com/tyomhk2015/gocoin/blockchain"

func main() {
	blockchain.Blockchain().AddBlock("FIRST")
	blockchain.Blockchain().AddBlock("SECOND")
	blockchain.Blockchain().AddBlock("THIRD")
	blockchain.Blockchain().AddBlock("FOURTH")
	blockchain.Blockchain().AddBlock("FIFTH")
}
