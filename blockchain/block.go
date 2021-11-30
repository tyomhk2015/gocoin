package blockchain

import (
	"crypto/sha256"
	"errors"
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

var BlockNotFoundErr = errors.New("The block you are looking for is not found.")

// FindBlock :Find one specific block.
func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, BlockNotFoundErr
	}
	block := &Block{}
	block.restore(blockBytes) // Load a block using blockByte(Hash).
	return block, nil
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}
