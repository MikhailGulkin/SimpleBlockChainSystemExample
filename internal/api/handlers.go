package api

import (
	"encoding/json"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api/static"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
	"net/http"
	"strconv"
)

type Handlers struct {
	blockChain *blockchain.BlockChain
}

func NewHandlers(bc *blockchain.BlockChain) *Handlers {
	return &Handlers{blockChain: bc}
}

func (h *Handlers) transactionFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(static.CreateTxForm))
	}
}

func (h *Handlers) processTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req TransactionRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Ошибка при чтении JSON", http.StatusBadRequest)
			return
		}
		amount, err := strconv.Atoi(req.Amount)
		if err != nil {
			http.Error(w, "Ошибка при чтении JSON", http.StatusBadRequest)
			return
		}
		txId := h.blockChain.PerformTransaction(req.Sender, req.Receiver, int64(amount))

		res := TransactionResponse{
			TransactionId: txId,
			FromAddress:   req.Sender,
			ToAddress:     req.Receiver,
			Amount:        int64(amount),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
func (h *Handlers) checkTransactionStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req TransactionStatusRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Ошибка при чтении JSON", http.StatusBadRequest)
			return
		}
		status := h.blockChain.CheckTransactionCompletion(req.TransactionId)
		res := TransactionStatusResponse{
			Status: status,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
func (h *Handlers) mineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req MineRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Ошибка при чтении JSON", http.StatusBadRequest)
			return
		}
		res := MineResponse{
			Message: "Транзакции успешно добавлены в блок",
		}
		err := h.blockChain.ProcessPendingTransaction(req.Address)
		if err != nil {
			res.Message = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
func (h *Handlers) getWallets(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		res := WalletResponse{
			Wallets: h.blockChain.GetWallets(),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
