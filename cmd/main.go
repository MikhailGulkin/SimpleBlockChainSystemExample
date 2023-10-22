package main

import (
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/wallet"
)

func main() {
	//wallets := wallet.NewWallets()
	//bc := blockchain.NewBlockChain(wallets.NewWallet().Address)
	//wallets.Save()
	//bc.Save()

	var bc blockchain.BlockChain
	var wallets wallet.Wallets
	bc.Load()
	wallets.Load()

	handlers := api.NewHandlers(&wallets, &bc)
	server := api.NewServer(handlers)
	server.SetupRoutes()
	server.Run()
}
