package blockchain

import (
	"fmt"
	"strings"
)

type Transaction struct {
	Id          string `json:"id"`
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Amount      int64  `json:"amount"`
}

func NewTransaction(fromAddress, toAddress string, amount int64) Transaction {
	return Transaction{
		Id:          GenerateTransactionId(),
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      amount,
	}
}
func (t *Transaction) ToString() string {
	return fmt.Sprintf("%s-%s-%s-%d", t.Id, t.FromAddress, t.ToAddress, t.Amount)
}

func TransactionsToString(transaction []Transaction) string {
	var str strings.Builder
	for _, tx := range transaction {
		str.WriteString(tx.ToString())
	}
	return str.String()
}
