package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func PoWSample() {
	// PoW and verification.
	difficulty := 3
	target := strings.Repeat("0", difficulty)
	nonce := 1
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("Proof_of_Work Verification"+fmt.Sprint(nonce))))
		fmt.Printf("\nCurrent Hash: %s\nTarget(Difficulty): %s\nNonce: %d\n", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			return
		}
		nonce++
	}
}
