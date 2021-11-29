package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/tyomhk2015/gocoin/db"
	"github.com/tyomhk2015/gocoin/utils"
)

type Block struct {
	Data         string `json:"data"`
	Hash         string `json:"hash"`
	PreviousHash string `json:"previoushash,omitempty"`
	Height       int    `json:"height"`
}

func createBlock(data string, previouHash string, height int) *Block {
	block := Block{
		Data:         data,
		Hash:         "",
		PreviousHash: previouHash,
		Height:       height,
	}
	// Save the following data in the database
	// Ingredients for hashes.
	payload := block.Data + block.PreviousHash + fmt.Sprint(block.Height)
	// Hash with the given data.
	block.Hash = fmt.Sprintf("%x", sha256.Sum224([]byte(payload)))
	// Until here
	block.persist()
	return &block
}

func (b *Block) persist() {
	// A function that saves the block in the database
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}
