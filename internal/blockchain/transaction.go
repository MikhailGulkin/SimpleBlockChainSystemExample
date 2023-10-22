package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"time"
)

type Transaction struct {
	Id              string                `json:"id"`
	FromAddress     string                `json:"fromAddress"`
	Time            string                `json:"time"`
	ToAddress       string                `json:"toAddress"`
	TransactionType utils.TransactionType `json:"transactionType"`
	Amount          int64                 `json:"amount"`
}

func NewTransaction(fromAddress, toAddress string, transactionType utils.TransactionType, amount int64) *Transaction {
	return &Transaction{
		Id:              utils.GenerateId(),
		Time:            time.Now().Format("2006.01.02 15:04:05"),
		FromAddress:     fromAddress,
		ToAddress:       toAddress,
		TransactionType: transactionType,
		Amount:          amount,
	}
}

func (bc *BlockChain) AddTransaction(
	fromAddress string,
	toAddress string,
	amount int64,
	transactionType utils.TransactionType,
	senderPublicKey *ecdsa.PublicKey,
	signature *utils.Signature,
) (string, error) {
	t := NewTransaction(fromAddress, toAddress, transactionType, amount)
	if fromAddress == MiningSender || fromAddress == bc.BlockChainAddress {
		bc.PendingTransactions = append(bc.PendingTransactions, t)
		return t.Id, nil
	}
	if !bc.VerifyTransactionSignature(senderPublicKey, signature, t) {
		return "", errors.New("invalid signature")
	}
	balance, err := bc.GetBalance(fromAddress)
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}
	if balance < amount {
		return "", errors.New("not enough balance")
	}
	bc.PendingTransactions = append(bc.PendingTransactions, t)
	return t.Id, nil
}

func (t *Transaction) ToSigJson() ([]byte, error) {
	return json.Marshal(struct {
		FromAddress     string                `json:"fromAddress"`
		ToAddress       string                `json:"toAddress"`
		TransactionType utils.TransactionType `json:"transactionType"`
		Amount          int64                 `json:"amount"`
	}{
		FromAddress:     t.FromAddress,
		ToAddress:       t.ToAddress,
		TransactionType: t.TransactionType,
		Amount:          t.Amount,
	})
}

func (bc *BlockChain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature,
	t *Transaction,
) bool {
	m, _ := t.ToSigJson()
	h := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}
