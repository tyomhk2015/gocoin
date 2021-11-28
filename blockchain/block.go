package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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
	db.SaveBlock(b.Hash, b.toBytes())
}

func (b *Block) toBytes() []byte {
	// Turn the data of block into bytes.
	// 'gob' package is used. https://pkg.go.dev/encoding/gob
	// 'gob' encode/decode bytes.
	var blockBuffer bytes.Buffer            // Set writer
	encoder := gob.NewEncoder(&blockBuffer) // Set the encoder with the writer.
	utils.HandleErr(encoder.Encode(b))      // Encode the block with the encoder and return the bytes to blockBuffer.
	return blockBuffer.Bytes()
}
