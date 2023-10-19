package main

import (
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain(make(map[string]int64))
	defer bc.Save()
	bc.Load()

	handlers := api.NewHandlers(&bc)
	server := api.NewServer(handlers)
	server.SetupRoutes()
	server.Run()
	
}
