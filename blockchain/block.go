package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tyomhk2015/gocoin/db"
	"github.com/tyomhk2015/gocoin/utils"
)

type Block struct {
	Hash         string `json:"hash"`
	PreviousHash string `json:"previoushash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	TimeStamp    int    `json:"timestamp"` // The time when the block is created.
	Transactions []*Tx  `json:"transactions"`
}

func createBlock(previouHash string, height int) *Block {
	block := Block{
		Hash:         "",
		PreviousHash: previouHash,
		Height:       height,
		Difficulty:   Blockchain().SetDifficulty(),
		Nonce:        0,
	}
	block.Transactions = Mempool.TxToConfirm()
	block.mine()
	block.persist()
	return &block
}

// PoW and verification.
func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.TimeStamp = int(time.Now().Unix()) // Save the time when the block is created.
		hash := utils.Hash(b)
		fmt.Printf("\n\n\nCurrent Hash: %s\nTarget(Difficulty): %s\nNonce: %d\n\n\n", hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		}
		b.Nonce++
	}
}

// /PoW and verification.

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
