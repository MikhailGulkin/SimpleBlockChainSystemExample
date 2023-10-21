package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
)

type Transaction struct {
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Amount      int64  `json:"amount"`
}

func NewTransaction(fromAddress, toAddress string, amount int64) *Transaction {
	return &Transaction{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      amount,
	}
}

func (bc *BlockChain) AddTransaction(
	fromAddress string,
	toAddress string,
	amount int64,
	senderPublicKey *ecdsa.PublicKey,
	signature *utils.Signature,
) (bool, error) {
	t := NewTransaction(fromAddress, toAddress, amount)
	if fromAddress == MiningSender {
		bc.PendingTransactions = append(bc.PendingTransactions, t)
		return true, nil
	}
	if !bc.VerifyTransactionSignature(senderPublicKey, signature, t) {
		return false, errors.New("invalid signature")
	}
	balance, err := bc.GetBalance(fromAddress)
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}
	if balance < amount {
		return false, errors.New("not enough balance")
	}
	bc.PendingTransactions = append(bc.PendingTransactions, t)
	return true, nil
}

func (bc *BlockChain) VerifyTransactionSignature(
	senderPublicKey *ecdsa.PublicKey,
	s *utils.Signature,
	t *Transaction,
) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}
