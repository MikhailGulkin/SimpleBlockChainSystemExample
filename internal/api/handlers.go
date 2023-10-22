package api

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api/static"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/wallet"
	"net/http"
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
		w.Write([]byte(static.CreateTxForm))
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
			http.Error(w, "Ошибка при чтении JSON", http.StatusBadRequest)
			return
		}
		tx := wallet.NewTransaction(
			user1.Address,
			user2.Address,
			utils.UserTransaction,
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
