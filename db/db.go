package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/tyomhk2015/gocoin/utils"
)

// Singletone Pattern of Database.
var db *bolt.DB

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

// Initialize a database.
func DB() *bolt.DB {
	// Open the my.db data file in the current directory.
	// It will be created if it doesn't exist.
	// 0600: Filemode, or chmod. https://en.wikipedia.org/wiki/Chmod
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		// bolt has concepts called buckets, where data is stored and organized.
		// E.g: bolt's bucket, similar to tables in sql.
		// Bucket one: Store data about blocks.
		// Bucket two: Store data about chains.
		// Create a bucket, https://github.com/boltdb/bolt#read-write-transactions
		err = db.Update(func(t *bolt.Tx) error {
			_, err = t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			utils.HandleErr(err)
			return err
		})
	}
	return db
}

// Save a block in the database.
// hash: key
// data: value
func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %x\n", hash, data)
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		fmt.Println("key:", []byte(hash), "\n\ndata:", data)
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

// Save the trace of the block in the chain and the database.
func SaveBlockchain(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte("checkpoint"), data)
		return err
	})
	utils.HandleErr(err)
}
