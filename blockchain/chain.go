package blockchain

import (
	"sync"

	"github.com/tyomhk2015/gocoin/db"
	"github.com/tyomhk2015/gocoin/utils"
)

type blockchain struct {
	// PoV: When the first user uses the blockchain for the first time.
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

// Can be used for only in the 'blockchain' package.
var b *blockchain

// A package for controlling sync of multithreads.
var once sync.Once

func Blockchain() *blockchain {
	if b == nil {
		// When the blockchain is request and if the blockchain is not initialized before
		// or requested for the first time, create a new blockchain.
		once.Do(func() {
			// Run this part only once, even though
			// this has been called several times by goroutine.
			// PoV: When the blockchain is created for the first time,
			//    : or the user is using the blockchain for the first time.
			b = &blockchain{"", 0}

			// TODO: Search for the checkpoint on the database.

			// TODO: If there is checkpoint, then restore the previous blockchain from bytes, stored in the DB.

			b.AddBlock("YAGOO")
		})
	}
	return b
}

func (b *blockchain) AddBlock(data string) {
	// Save the block data to database.
	block := createBlock(data, b.NewestHash, b.Height+1)

	// Renew the chain data.
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
