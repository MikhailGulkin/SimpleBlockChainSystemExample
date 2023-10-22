package api

type (
	TransactionRequest struct {
		Sender   string `json:"sender,omitempty"`
		Receiver string `json:"receiver,omitempty"`
		Amount   string `json:"amount,omitempty"`
	}
	MineRequest struct {
		Address string `json:"address,omitempty"`
	}
	TransactionStatusRequest struct {
		TransactionId string `json:"transactionId,omitempty"`
	}
	TransactionResponse struct {
		TransactionId string `json:"transactionId"`
		FromAddress   string `json:"fromAddress"`
		ToAddress     string `json:"toAddress"`
		Amount        int64  `json:"amount"`
	}
	TransactionStatusResponse struct {
		Status bool `json:"status"`
	}
	MineResponse struct {
		Message string `json:"message"`
	}

	WalletsResponse struct {
		Wallets map[string]int64 `json:"wallets"`
	}
	WalletResponse struct {
		Wallet string `json:"wallet"`
	}
	CreateWalletResponse struct {
		Address string `json:"address"`
	}
	BlockChainValidityResponse struct {
		IsValid bool `json:"isValid"`
	}
)
