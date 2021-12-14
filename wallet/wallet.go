package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hashWalletFile() bool {
	_, err := os.Stat("KRONI")
	return !os.IsNotExist(err)
}

func Wallet() *wallet {
	if w == nil {
		if hashWalletFile() {
			// Restore wallet from the file.
			fmt.Println("EXIST")
		}
		// Else, create Private Key and save to as a file
		fmt.Println("DO NOT EXIST")
	}
	return w
}
