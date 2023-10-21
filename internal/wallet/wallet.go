package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"log"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey `json:"privateKey"`
	PublicKey  *ecdsa.PublicKey  `json:"publicKey"`
	Address    string            `json:"address"`
}

func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.PrivateKey = privateKey
	w.PublicKey = &w.PrivateKey.PublicKey

	w.Address = generateAddressFromKeys(w.PublicKey)
	return w
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
