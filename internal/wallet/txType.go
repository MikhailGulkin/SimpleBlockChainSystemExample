package wallet

import "github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"

func GetTransactionType(amount int64) utils.TransactionType {
	if amount <= 10 {
		return utils.UserTransactionMin
	} else if amount > 10 && amount <= 100 {
		return utils.UserTransactionAvg
	}
	return utils.UserTransactionMax
}
