package blockchain

import (
	"errors"
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

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

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

func makeTx(giver, receiver string, amount int) (*Tx, error) {
	// Check if the user has enough coins for the transaction.
	if Blockchain().BalanceByAddress(giver) < amount {
		return nil, errors.New("Not enough coins.")
	}

	// Make as many Txs as possible until the 'amount' is met.
	var txIns []*TxIn
	var txOuts []*TxOut // For giving some changes back to giver, the remainder after the 'amount'.
	totalInput := 0
	oldTxOuts := Blockchain().TxOutsByAddress(giver)
	for _, txOut := range oldTxOuts {
		if totalInput >= amount {
			// If totalInput's accumulation satisfies amount
			break
		}
		txIn := &TxIn{txOut.Owner, txOut.Balance}
		txIns = append(txIns, txIn)
		totalInput += txOut.Balance
	}

	// Check if there is any remainder that the giver has to get back.
	change := totalInput - amount
	if change > 0 {
		// Make a new transaction for reclaiming remainder.
		changeTxOut := &TxOut{giver, change}
		txOuts = append(txOuts, changeTxOut)
	}

	// Create and add output transaction of 'quotient'.
	txOut := &TxOut{receiver, amount}
	txOuts = append(txOuts, txOut)

	// Make a transaction to add to the mempool.
	tx := &Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil
}

func (m *mempool) AddTx(receiver string, amount int) error {
	tx, err := makeTx("HOLOLIVE", receiver, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

// Get transactions from the mempool, and erase the confirmed transactions.
func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("HOLOLIVE")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
