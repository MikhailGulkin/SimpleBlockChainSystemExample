package api

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/wallet"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Handlers struct {
	wallets    *wallet.Wallets
	blockChain *blockchain.BlockChain
}

func NewHandlers(wallets *wallet.Wallets, bc *blockchain.BlockChain) *Handlers {
	return &Handlers{blockChain: bc, wallets: wallets}
}

func (h *Handlers) transactionFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")

		file, err := os.Open("internal/api/static/createTransaction.html")
		defer file.Close()

		if err != nil {
			http.Error(w, "Не удалось открыть файл", http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Не удалось отправить файл", http.StatusInternalServerError)
		}
	}
}

func (h *Handlers) processTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")

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
		user1 := h.wallets.GetWallet(req.Sender)
		user2 := h.wallets.GetWallet(req.Receiver)
		if user1 == nil || user2 == nil {
			http.Error(w, "Один из пользователей временно не доступен или не существует", http.StatusBadRequest)
			return
		}
		tx := wallet.NewTransaction(
			user1.Address,
			user2.Address,
			wallet.GetTransactionType(int64(amount)),
			user1.PrivateKey,
			user1.PublicKey,
			int64(amount),
		)
		sig := tx.GenerateSignature()
		transactionId, err := h.blockChain.AddTransaction(
			tx.FromAddress,
			tx.ToAddress,
			tx.Amount,
			tx.TransactionType,
			tx.PublicKey,
			sig,
		)
		if err != nil {
			http.Error(w, fmt.Errorf("ошибка при создании транзакции %w", err).Error(), http.StatusBadRequest)
			return
		}

		res := TransactionResponse{
			TransactionId: transactionId,
			FromAddress:   req.Sender,
			ToAddress:     req.Receiver,
			Amount:        int64(amount),
		}

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
		err := h.blockChain.Mining(req.Address)
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
		res := WalletsResponse{
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
func (h *Handlers) checkBlockChainValidity(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		res := BlockChainValidityResponse{
			IsValid: h.blockChain.ValidChain(h.blockChain.GetChain()),
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
func (h *Handlers) createWallet(writer http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		writer.Header().Set("Content-Type", "application/json")

		w := h.wallets.NewWallet()
		h.blockChain.RegisterNewWallet(w.Address)
		res := CreateWalletResponse{
			Address: w.Address,
		}
		if err := json.NewEncoder(writer).Encode(res); err != nil {
			http.Error(writer, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(writer, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
func (h *Handlers) getWalletTransactions(writer http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		writer.Header().Set("Content-Type", "application/json")

		address := r.URL.Query().Get("address")
		if address == "" {
			http.Error(writer, "Не указан адрес", http.StatusBadRequest)
			return
		}
		user := h.wallets.GetWallet(address)
		if user == nil {
			http.Error(writer, "Пользователь не найден", http.StatusBadRequest)
			return
		}
		transactions := h.blockChain.GetWalletTransactions(user.Address)
		res := AllWalletTransactionsResponse{
			Transactions: transactions,
		}
		if err := json.NewEncoder(writer).Encode(res); err != nil {
			http.Error(writer, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(writer, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) getGenBlocks(writer http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		blocks := h.blockChain.GenerateBlocks(10)
		res := make([][]TransactionGenBlocksResponse, 10)
		for index, block := range blocks {
			var txs []TransactionGenBlocksResponse
			for _, tx := range block.Transactions {
				txs = append(txs, TransactionGenBlocksResponse{
					string(tx.TransactionType),
				})
			}
			res[index] = txs
		}
		if err := json.NewEncoder(writer).Encode(res); err != nil {
			http.Error(writer, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(writer, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
