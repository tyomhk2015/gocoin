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
	Difficulty int    `json:"difficulty"`
}

// Can be used for only in the 'blockchain' package.
var b *blockchain

// A package for controlling sync of multithreads.
var once sync.Once

func Blockchain() *blockchain {
	// if b == nil {
	// 	// When the blockchain is request and if the blockchain is not initialized before
	// 	// or requested for the first time, create a new blockchain.
	// 	once.Do(func() {
	// 		// Run this part only once, even though
	// 		// this has been called several times by goroutine.
	// 		// PoV: When the blockchain is created for the first time,
	// 		//    : or the user is using the blockchain for the first time.
	// 		b = &blockchain{Height: 0}

	// 		// Search for the checkpoint on the database,
	// 		// meaning this checks if previous blockchain is in the DB.
	// 		checkPoint := db.CheckPoint()
	// 		// fmt.Printf("\nBEFORE\nNewHash: %x\nHeight: %d\nCheckpoint: %x", b.NewestHash, b.Height, checkPoint)
	// 		if checkPoint == nil {
	// 			// If there is no exisiting blockchain, create a new one.
	// 			b.AddBlock()
	// 		} else {
	// 			// If there is checkpoint, then restore the previous blockchain from bytes, stored in the DB.
	// 			fmt.Println("\n\n**RESTORING THE BLOCKCHAIN FROM DB.**")
	// 			b.restoreBlockChain(checkPoint)
	// 		}
	// 	})
	// 	// fmt.Printf("\nAFTER\nNewHash: %x\nHeight: %d\n", b.NewestHash, b.Height)
	// }
	// // fmt.Println(b.NewestHash)
	// return b

	once.Do(func() {
		b = &blockchain{Height: 0}
		checkPoint := db.CheckPoint()
		if checkPoint == nil {
			b.AddBlock()
		} else {
			fmt.Println("\n\n**RESTORING THE BLOCKCHAIN FROM DB.**")
			b.restoreBlockChain(checkPoint)
		}
	})
	return b
}

func (b *blockchain) AddBlock() {
	// Save the block data to database.
	block := createBlock(b.NewestHash, b.Height+1, SetDifficulty(b))

	// Renew the chain data.
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.Difficulty = block.Difficulty
	persistBlockchain(b)
}

// Save the blockchain in bytes.
func persistBlockchain(b *blockchain) {
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
func Blocks(b *blockchain) []*Block {
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

// PoW, setting difficulty.
const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5 // Checkpoint of deciding wether the difficulty should be harder/easier/ stay the same.
	intervalPerBlock   int = 1 // The unit is in minutes.
	timeFlexibility    int = 2
)

func SetDifficulty(b *blockchain) int {
	if b.Height == 0 {
		// Set difficulty to 2, when the blockchain is newly created.
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// Renew the difficulty.
		return renewDifficulty(b)
	}
	return b.Difficulty
}

func renewDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
	newestBlock := allBlocks[0] // The blocks are in descending order, the newest one is at the start.
	lastDifficultyBlock := allBlocks[difficultyInterval-1]
	elapsedTime := (newestBlock.TimeStamp - lastDifficultyBlock.TimeStamp) / 60 // Calculate in minutes.
	expectedTime := difficultyInterval * intervalPerBlock

	if elapsedTime <= (expectedTime - timeFlexibility) {
		return b.Difficulty + 1
	} else if elapsedTime >= (expectedTime + timeFlexibility) {
		return b.Difficulty - 1
	}
	return b.Difficulty
}

// /PoW, setting difficulty.

// Get all the transaction Outputs from each block.
func (b *blockchain) txOuts() []*TxOut {
	var txOuts []*TxOut
	blocks := Blocks(b)
	for _, block := range blocks {
		for _, tx := range block.Transactions {
			txOuts = append(txOuts, tx.TxOuts...)
		}
	}
	return txOuts
}

// Get all the 'unspent' transaction outputs of a specific address or owner.
func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var unspentTxOuts []*UTxOut
	txsReferredAtInput := make(map[string]bool)
	for _, block := range Blocks(b) {
		for _, tx := range block.Transactions {
			for _, txIn := range tx.TxIns {
				// Collect IDs of the transactions that were referred at INPUT,
				// of the specifed user or wallet.
				if txIn.Owner == address {
					txsReferredAtInput[txIn.TxID] = true // A flag, meaning this Tx has been 'referred' at INPUT.
				}
			}
			for index, txOut := range tx.TxOuts {
				// Collect transactions that were not referred at INPUT,
				// of the specifed user or wallet.
				if txOut.Owner == address {
					_, referred := txsReferredAtInput[tx.ID]
					if !referred {
						// Check if UTx are used in mempool.
						uTxOut := &UTxOut{tx.ID, index, txOut.Balance}
						if !isOnMempool(uTxOut) {
							// Add UTx that are not used in mempool.
							unspentTxOuts = append(unspentTxOuts, &UTxOut{tx.ID, index, txOut.Balance})
						}
					}
				}
			}
		}
	}
	return unspentTxOuts
}

// Get total balance of 'unspent' transaction outputs from a specific address.
func BalanceByAddress(address string, b *blockchain) int {
	var totalBalance int
	txOuts := UTxOutsByAddress(address, b)
	for _, txOut := range txOuts {
		totalBalance += txOut.Balance
	}
	return totalBalance
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
Outer: // Label, use for stopping nested for loops. This can break loops anywhere.
	for _, tx := range Mempool.Txs {
		for _, txIn := range tx.TxIns {
			if txIn.TxID == uTxOut.TxID && txIn.Index == uTxOut.Index {
				exists = true
				break Outer
			}
		}
	}
	return exists
}
