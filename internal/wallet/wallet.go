package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string `json:"address"`
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.PrivateKey = privateKey
	w.PublicKey = &w.PrivateKey.PublicKey

	w.Address = generateAddressFromKeys(w.PublicKey)
	return w
}
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.PrivateKey.D.Bytes())
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%064x%064x", w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes())
}

func (w *Wallet) ToJson() ([]byte, error) {
	return json.Marshal(struct {
		PrivateKey        string `json:"privateKey"`
		PublicKey         string `json:"publicKey"`
		BlockchainAddress string `json:"blockchainAddress"`
	}{
		PrivateKey:        w.PrivateKeyStr(),
		PublicKey:         w.PublicKeyStr(),
		BlockchainAddress: w.Address,
	})
}

func (w *Wallet) MarshalJSON() ([]byte, error) {
	privateKeyBytes := w.PrivateKey.D.Bytes()
	publicKeyXBytes := w.PublicKey.X.Bytes()
	publicKeyYBytes := w.PublicKey.Y.Bytes()

	type Alias Wallet
	return json.Marshal(&struct {
		Address    string `json:"address"`
		PrivateKey []byte `json:"privateKey"`
		PublicKeyX []byte `json:"publicKeyX"`
		PublicKeyY []byte `json:"publicKeyY"`
	}{
		Address:    w.Address,
		PrivateKey: privateKeyBytes,
		PublicKeyX: publicKeyXBytes,
		PublicKeyY: publicKeyYBytes,
	})
}

func (w *Wallet) UnmarshalJSON(data []byte) error {
	type Alias Wallet
	aux := &struct {
		Address    string `json:"address"`
		PrivateKey []byte `json:"privateKey"`
		PublicKeyX []byte `json:"publicKeyX"`
		PublicKeyY []byte `json:"publicKeyY"`
	}{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	w.Address = aux.Address
	w.PrivateKey = new(ecdsa.PrivateKey)
	w.PrivateKey.PublicKey.Curve = elliptic.P256()
	w.PrivateKey.D = new(big.Int).SetBytes(aux.PrivateKey)

	w.PublicKey = new(ecdsa.PublicKey)
	w.PublicKey.Curve = elliptic.P256()
	w.PublicKey.X = new(big.Int).SetBytes(aux.PublicKeyX)
	w.PublicKey.Y = new(big.Int).SetBytes(aux.PublicKeyY)

	return nil
}
