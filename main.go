package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	tempChain := blockchain{}
	tempChain.appendBlock("First Block")
	tempChain.appendBlock("Second Block")
	tempChain.appendBlock("Third Block")
	tempChain.appendBlock("Fourth Block")
	tempChain.showBlockchain()
}

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) appendBlock(data string) {
	// data: The data the block holds
	// hash: A tag or an address of the block, created with hash function using 'data' and 'previous hash' string.
	//       Most of hashes are written in hexdecimal, 16bit 0~F (%x).
	// prevHash: A tag or an address of the previous block, refers to the 'hash' of previous block.
	//           If this is the initial block, prevHash will have an empty string, "".
	createdBlock := block{data, "", ""}
	createdBlock.prevHash = b.getPreviousHash()
	createdBlock.hash = getHash(data, createdBlock.prevHash)
	b.blocks = append(b.blocks, createdBlock)
}

func getHash(data, prevHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data+prevHash)))
}

func (b *blockchain) getPreviousHash() string {
	fmt.Println(b.blocks)
	if len(b.blocks) > 0 {
		return (b.blocks[len(b.blocks)-1].hash)
	}
	return ""
}

func (b *blockchain) showBlockchain() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\nHash: %s\nPrevious Hash: %s\n\n", block.data, block.hash, block.prevHash)
	}
}
