package blockchain

import (
	"fmt"
	"sync"
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
			b = &blockchain{"", 0}
			b.AddBlock("YAGOO")
		})
	}
	fmt.Println("SOMETHING IN B out", b)
	return b
}

func (b *blockchain) AddBlock(data string) {
	// Save the block data to database.
	block := createBlock(data, b.NewestHash, b.Height)

	// Renew the chain data.
	b.NewestHash = block.Hash
	b.Height = block.Height
}
