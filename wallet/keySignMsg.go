package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/tyomhk2015/gocoin/utils"
)

const (
	hashedMessage string = "6586f3af7a015f2884e5e651539b18efe35d587d896757161442b8068c1602ca" // Subsitute of transaction ID.
	signature     string = "3a4be3a3146db9f8aae3b185b1e3505a1393a8ca608354271e135115f3a29b9542468eb5d5e93383f5b92bb7f1043f37f29cc46b80cda5771249ad5ede297691"
	privateKey    string = "3077020101042003da0ad9abf4460301350004ca81eaccbd12fbdacf5c0b70f97a33a205497005a00a06082a8648ce3d030107a1440342000467b20e00a65ec10d51e79df88efe25ebafbb21520898edf41ae7c6645df3c72006d6e206398fc0395f7ec462bb10112c477a2098fe1e3ba73aba24c3376d49c8"
)

func restoreKeySign() {
	// Check if the loaded hashes are encrypted in hexadecimal.
	// If the hex code has been manipulated, an error will be returned.
	privateKeyBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	// Take private key bytes to restore the key
	restoredPrivateKey, err := x509.ParseECPrivateKey(privateKeyBytes)
	utils.HandleErr(err)

	// Signature is in [[32],[32]]
	// Restore the signature
	signatureByte, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	leftBytes := signatureByte[:len(signatureByte)/2]  // From 0 to designated index.
	rightBytes := signatureByte[len(signatureByte)/2:] // From designated index to the end.

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(leftBytes)
	bigS.SetBytes(rightBytes)

	recoveredHashByte, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	isVerified := ecdsa.Verify(&restoredPrivateKey.PublicKey, recoveredHashByte, &bigR, &bigS)
	fmt.Println(isVerified)
}

func createMsgKeySignVerification() {
	// Caution: Neither practical nor realistic

	// Generating Private key and Public key.
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // Has to be saved as a file.
	utils.HandleErr(err)
	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey) // Make private key recoverable.
	fmt.Printf("%x\n\n\n", keyAsBytes)
	utils.HandleErr(err)

	// Generate signature.
	message := "NIJISANZI + HOLOLIVE + VOMS" // Subsitute of transaction ID.
	hashedMessage := utils.Hash(message)
	fmt.Println(hashedMessage)
	HashMsgAsByte, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, HashMsgAsByte) // Signature
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("%x\n", signature)
	utils.HandleErr(err)

	// Verification
	isVerified := ecdsa.Verify(&privateKey.PublicKey, HashMsgAsByte, r, s)
	fmt.Println(isVerified)
}
