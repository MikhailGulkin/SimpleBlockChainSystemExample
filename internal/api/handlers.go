package api

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api/static"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
	"net/http"
	"strconv"
)

type Handlers struct {
	blockChain *blockchain.BlockChain
}

func transactionFormHandler(w http.ResponseWriter, r *http.Request) {
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
		h.blockChain.PerformTransaction(req.Sender, req.Receiver, int64(amount))

		res := TransactionResponse{
			Message: fmt.Sprintf("Транзакция в обработке: От кого: %s, Кому: %s, Сколько: %s",
				req.Sender, req.Receiver, req.Amount),
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
