package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
)

func generateAddressFromKeys(pubKey *ecdsa.PublicKey) string {
	pubKeyBytes := elliptic.Marshal(pubKey.Curve, pubKey.X, pubKey.Y)
	hash := sha256.Sum256(pubKeyBytes)
	return fmt.Sprintf("0x%x", hash[:20])
}
