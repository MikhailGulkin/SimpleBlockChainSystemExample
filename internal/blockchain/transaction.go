package blockchain

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	id          string
	fromAddress string
	toAddress   string
	amount      int64
}

func NewTransaction(fromAddress, toAddress string, amount int64) Transaction {
	return Transaction{
		id:          GenerateTransactionId(),
		fromAddress: fromAddress,
		toAddress:   toAddress,
		amount:      amount,
	}
}
func (t *Transaction) ToString() string {
	return fmt.Sprintf("%s-%s-%s-%d", t.id, t.fromAddress, t.toAddress, t.amount)
}

func TransactionsToString(transaction []Transaction) string {
	var str strings.Builder
	for _, tx := range transaction {
		str.WriteString(tx.ToString())
	}
	return str.String()
}
func GenerateTransactionId() string {
	var result strings.Builder
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		if i%4 == 0 && i != 0 {
			result.WriteString("-")
		}
		if i%2 == 0 {
			result.WriteString(strconv.Itoa(random.Intn(10)))
		} else {
			result.WriteString(string(rune(random.Intn(26) + 65)))
		}
	}
	return result.String()
}
