package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/tyomhk2015/gocoin/utils"
)

const (
	hashedMessage string = "6586f3af7a015f2884e5e651539b18efe35d587d896757161442b8068c1602ca" // Subsitute of transaction ID.
	signature     string = "67e32ce67cc0607c788ab11fb3df4d2affe1464ba32be8166653c666e447e91e7e770ec00ce36ee825c19e4af1f14ffbdca28962e801620910d697e5c22f0d43"
	privateKey    string = "3077020101042096f6f2dc95081a755b461b20b1419a205c1704354957b86ceba6ae8a3121aa84a00a06082a8648ce3d030107a1440342000443cfe2873a3ab81c5fe1972d72986570e631fed225c4413f49f13d943be20ea3740ac1098e0c60abd7f6384207a769886cb489e452df1707b21bbd076c0f67c6"
)

func Start() {

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
