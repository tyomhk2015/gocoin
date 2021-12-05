package blockchain

import (
	"time"

	"github.com/tyomhk2015/gocoin/utils"
)

// Transaction struct
type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

type TxOut struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

const (
	minerReward int = 7
)

func makeCoinbaseTx(address string /* miner's address */) *Tx {
	// Some coins are about to be sent from 'Metamask', the miner.
	txIns := []*TxIn{
		{"Metamask", minerReward},
	}
	// Some coins are about to be comming from 'Metamask' to designated miner.
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}
