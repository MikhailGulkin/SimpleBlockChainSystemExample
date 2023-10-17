package api

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api/static"
	"net/http"
)

func transactionFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Чтение содержимого .html файла

		// Отправка содержимого файла в ответе
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(static.CreateTxForm))
	}
}

type TransactionRequest struct {
	Sender   string `json:"sender,omitempty"`
	Receiver string `json:"receiver,omitempty"`
	Amount   string `json:"amount,omitempty"`
}
type TransactionResponse struct {
	Message string `json:"message"`
}

func processTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req TransactionRequest

		// Раскодировать JSON-запрос в структуру TransactionRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Ошибка при чтении JSON", http.StatusBadRequest)
			return
		}

		// Здесь вы можете обработать данные, например, выполнить транзакцию или другие действия
		fmt.Println(req)
		// Создать JSON-ответ
		res := TransactionResponse{Message: fmt.Sprintf("Транзакция в обработке: От кого: %s, Кому: %s, Сколько: %s",
			req.Sender, req.Receiver, req.Amount),
		}

		// Сериализовать JSON-ответ
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Ошибка при отправке JSON-ответа", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
