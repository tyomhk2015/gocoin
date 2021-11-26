// Singleton pattern:
// Always sharing ONLY one instance throughout the whole program.
// The instance can be accessed anywhere in program, global access.
// ?: Referring to the same instance means that once the singleton instance
//    is created, this memory address will be fixed like 'static' of Java, I guess.
// Ref: https://refactoring.guru/design-patterns/singleton

package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data         string `json:"data"`
	Hash         string `json:"hash"`
	PreviousHash string `json:"previoushash,omitempty"`
	Height       int    `json:"height"`
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
	newBlock := Block{data, "", getPreviousHash(), len(GetBlockchain().blocks) + 1}
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
	// Get all the blocks.
	return b.blocks
}

var NotFoundErr = errors.New("The block does not exist.")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	// Get only one block.
	if height > len(b.blocks) {
		// If the requested 'height' exceeds the length of the blockchain, return error message.
		return nil, NotFoundErr
	}
	return b.blocks[height-1], nil
}
