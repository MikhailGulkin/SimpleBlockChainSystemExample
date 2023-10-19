package api

import (
	"log"
	"net/http"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{handlers: handlers}
}

func (h *Server) SetupRoutes() {
	http.HandleFunc("/", h.handlers.transactionFormHandler)
	http.HandleFunc("/process-transaction", h.handlers.processTransaction)
	http.HandleFunc("/mine", h.handlers.mineHandler)
	http.HandleFunc("/check-tx-status", h.handlers.checkTransactionStatus)
	http.HandleFunc("/get-wallets", h.handlers.getWallets)
}

func (h *Server) Run() {
	log.Printf("server is running on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func disableCors(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	h := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler(w, r)
	}
	return h

}
