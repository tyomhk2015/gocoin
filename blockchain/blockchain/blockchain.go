// Singleton pattern:
// Always sharing ONLY one instance throughout the whole program.
// The instance can be accessed anywhere in program, global access.
// ?: Referring to the same instance means that once the singleton instance
//    is created, this memory address will be fixed like 'static' of Java, I guess.
// Ref: https://refactoring.guru/design-patterns/singleton

package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data         string
	Hash         string
	PreviousHash string
}

type blockchain struct {
	blocks []*Block
}

// Can be used for only in the 'blockchain' package.
var b *blockchain

// A package for controlling sync of multithreads.
var once sync.Once

func GetBlockchain() *blockchain {
	if b == nil {
		// When the blockchain is request and if the blockchain is not initialized before
		// or requested for the first time, create a new blockchain.
		once.Do(func() {
			// Run this part only once, even though
			// this has been called several times by goroutine.
			b = &blockchain{}
			b.AddBlock("YAGOO")
		})
	}
	return b
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getPreviousHash()}
	newBlock.Hash = newBlock.createHash()
	return &newBlock
}

func (b *Block) createHash() string {
	hash := sha256.Sum256([]byte(b.Data + b.PreviousHash))
	return fmt.Sprintf("%x", hash)
}

func getPreviousHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b *blockchain) ShowAllBlocks() []*Block {
	return b.blocks
}
