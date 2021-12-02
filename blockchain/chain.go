package blockchain

import (
	"fmt"
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

			// Search for the checkpoint on the database,
			// meaning this checks if previous blockchain is in the DB.
			checkPoint := db.CheckPoint()
			fmt.Printf("\nBEFORE\nNewHash: %x\nHeight: %d\nCheckpoint: %x", b.NewestHash, b.Height, checkPoint)
			if checkPoint == nil {
				// If there is no exisiting blockchain, create a new one.
				b.AddBlock("YAGOO")
			} else {
				// If there is checkpoint, then restore the previous blockchain from bytes, stored in the DB.
				fmt.Println("\n\n**RESTORING THE BLOCKCHAIN FROM DB.**")
				b.restoreBlockChain(checkPoint)
			}
		})
		fmt.Printf("\nAFTER\nNewHash: %x\nHeight: %d\n", b.NewestHash, b.Height)
	}
	fmt.Println(b.NewestHash)
	return b
}

func (b *blockchain) AddBlock(data string) {
	// Save the block data to database.
	block := createBlock(data, b.NewestHash, b.Height+1)

	// Renew the chain data.
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// Save the blockchain in bytes.
func (b *blockchain) persist() {
	db.SaveCheckPoint(utils.ToBytes(b))
}

// Convert the blockchain from bytes to go lang, the data conversion.
// Read from database and change the pointer of the 'b' to the read blockchain data.
func (b *blockchain) restoreBlockChain(data []byte) {
	// Set a decoder with a target to decode.
	// decoder := gob.NewDecoder(bytes.NewReader(data))
	// Replace the memory address of 'b', with the memory address of loaded blockchain.
	// err := decoder.Decode(b)
	// utils.HandleErr(err)
	utils.FromBytes(b, data)
}

// Get all the blocks by using the recent block's previous hash
// and iterate the process until there are no more blocks left.
// Traverse from recent block to the first block on the blockchain.
func (b *blockchain) Blocks() []*Block {
	var retrievedBlocks []*Block // A temporary storage for collecting retrieved blocks.
	hashCursor := b.NewestHash
	for {
		retrievedBlock, _ := FindBlock(hashCursor)
		retrievedBlocks = append(retrievedBlocks, retrievedBlock)
		if retrievedBlock.PreviousHash != "" || len(retrievedBlock.PreviousHash) != 0 {
			hashCursor = retrievedBlock.PreviousHash
		} else {
			break // Without 'break' the for loop will continue forever.
		}
	}
	return retrievedBlocks
}
