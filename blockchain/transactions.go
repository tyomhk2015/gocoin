package blockchain

import (
	"errors"
	"time"

	"github.com/tyomhk2015/gocoin/utils"
)

// Transaction struct
type Tx struct {
	ID        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	TxID  string `json:"txID"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
}

// This structs is for finding 'unspent' transaction outputs.
type UTxOut struct {
	TxID    string `json:"txID"`
	Index   int    `json:"index"`
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
		{"", -1, "Binance"},
	}
	// Some coins are about to be comming from 'Metamask' to designated miner.
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

func (t *Tx) getId() {
	t.ID = utils.Hash(t)
}

// Make a transaction with 'unspent' transactions.
func makeTx(giver, receiver string, amount int) (*Tx, error) {
	// Check if the giver has enough coins by checking UTxs.
	if BalanceByAddress(giver, Blockchain()) < amount {
		return nil, errors.New("Not enough money!")
	}

	// Makes slices for inputs and outputs, a transaction.
	var txOuts []*TxOut
	var txIns []*TxIn
	totalBalance := 0
	uTxOuts := UTxOutsByAddress(giver, Blockchain())

	// Amend the required amount by accumulating balance of 'unspent' transactions.
	for _, uTxOut := range uTxOuts {
		if totalBalance >= amount {
			// If the accumulated balance satisfies required amount, quit the loop.
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, giver}
		txIns = append(txIns, txIn) // This 'unspent' transaction will be used/referred as INPUT of transaction.
		totalBalance += uTxOut.Balance
	}

	// Check if there is need for changes.
	// If changes are needed, make a transaction output for it.
	change := totalBalance - amount
	if change > 0 {
		changeTxOut := &TxOut{giver, change}
		txOuts = append(txOuts, changeTxOut)
	}

	// Add a transaction that meets the required amount of coins.
	txOut := &TxOut{receiver, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId() // Add an ID for this new transaction.
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
