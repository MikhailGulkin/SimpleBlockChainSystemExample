package utils

type TransactionType string

const (
	UserTransaction TransactionType = "USER_TRANSACTION"
	WalletCreate    TransactionType = "WALLET_CREATE"
	Mining          TransactionType = "MINING"
)
