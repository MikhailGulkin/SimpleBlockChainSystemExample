package api

type TransactionRequest struct {
	Sender   string `json:"sender,omitempty"`
	Receiver string `json:"receiver,omitempty"`
	Amount   string `json:"amount,omitempty"`
}
type MineRequest struct {
	Address string `json:"address,omitempty"`
}

type TransactionResponse struct {
	Message string `json:"message"`
}
