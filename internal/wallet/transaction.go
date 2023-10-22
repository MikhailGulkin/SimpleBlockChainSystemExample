package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"log"
)

type Transaction struct {
	FromAddress     string
	ToAddress       string
	TransactionType utils.TransactionType
	PrivateKey      *ecdsa.PrivateKey
	PublicKey       *ecdsa.PublicKey
	Amount          int64
}

func NewTransaction(
	fromAddress string,
	toAddress string,
	transactionType utils.TransactionType,
	privateKey *ecdsa.PrivateKey,
	publicKey *ecdsa.PublicKey,
	amount int64,
) *Transaction {
	return &Transaction{
		FromAddress:     fromAddress,
		ToAddress:       toAddress,
		TransactionType: transactionType,
		PrivateKey:      privateKey,
		PublicKey:       publicKey,
		Amount:          amount,
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
		FromAddress     string `json:"fromAddress"`
		ToAddress       string `json:"toAddress"`
		TransactionType string `json:"transactionType"`
		Amount          int64  `json:"amount"`
	}{
		FromAddress:     t.FromAddress,
		ToAddress:       t.ToAddress,
		TransactionType: string(t.TransactionType),
		Amount:          t.Amount,
	})
}
