package main

import (
	"github.com/tyomhk2015/gocoin/blockchain"
	"github.com/tyomhk2015/gocoin/cli"
	"github.com/tyomhk2015/gocoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
