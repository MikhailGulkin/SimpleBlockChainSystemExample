package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"log"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
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

func (w *Wallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PrivateKey        string `json:"PrivateKey"`
		PublicKey         string `json:"PublicKey"`
		BlockchainAddress string `json:"blockchainAddress"`
	}{
		PrivateKey:        w.PrivateKeyStr(),
		PublicKey:         w.PublicKeyStr(),
		BlockchainAddress: w.Address,
	})
}

type Transaction struct {
	FromAddress string
	ToAddress   string
	PrivateKey  *ecdsa.PrivateKey
	PublicKey   *ecdsa.PublicKey
	Amount      int64
}

func NewTransaction(
	fromAddress string,
	toAddress string,
	privateKey *ecdsa.PrivateKey,
	publicKey *ecdsa.PublicKey,
	amount int64,
) *Transaction {
	return &Transaction{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
		Amount:      amount,
	}
}

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)

	log.Println("Generate signature", string(m))

	h := sha256.Sum256(m)
	r, s, err := ecdsa.Sign(rand.Reader, t.PrivateKey, h[:])
	if err != nil {
		// TODO: handle error
		return nil
	}
	return &utils.Signature{R: r, S: s}
}
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FromAddress string `json:"fromAddress"`
		ToAddress   string `json:"toAddress"`
		Amount      int64  `json:"amount"`
	}{
		FromAddress: t.FromAddress,
		ToAddress:   t.ToAddress,
		Amount:      t.Amount,
	})
}
