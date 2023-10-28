package utils

type TransactionType string

const (
	UserTransactionMin TransactionType = "USER_TRANSACTION_MIN"
	UserTransactionAvg TransactionType = "USER_TRANSACTION_AVG"
	UserTransactionMax TransactionType = "USER_TRANSACTION_MAX"
	WalletCreate       TransactionType = "WALLET_CREATE"
	Mining             TransactionType = "MINING"
)
