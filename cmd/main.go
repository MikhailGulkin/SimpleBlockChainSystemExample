package main

import (
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
)

func main() {
	var bc blockchain.BlockChain
	bc.Load()
	fmt.Println(bc.ValidChain(bc.Chain))
	//bc := blockchain.NewBlockChain(make(map[string]int64))
	//defer bc.Save()
	//bc.Load()
	//
	//handlers := api.NewHandlers(&bc)
	//server := api.NewServer(handlers)
	//server.SetupRoutes()
	//server.Run()
}
