package api

import "net/http"

func Run() error {
	http.HandleFunc("/", transactionFormHandler)
	//http.HandleFunc("/process-transaction", Handlers{}.processTransaction)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return err
	}
	return nil
}
