package main

import (
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/wallet"
)

func main() {
	wallets := wallet.NewWallets()
	w := wallets.NewWallet()
	user1 := wallets.NewWallet()
	user2 := wallets.NewWallet()
	wallets.Save()
	bc := blockchain.NewBlockChain(w.Address)
	bc.RegisterNewWallet(user1.Address)
	bc.RegisterNewWallet(user2.Address)
	bc.AddTransaction(
		blockchain.MiningSender,
		user1.Address,
		100,
		nil,
		nil,
	)
	bc.AddTransaction(
		blockchain.MiningSender,
		user2.Address,
		100,
		nil,
		nil,
	)
	bc.Mining()
	fmt.Println(bc.GetWallets())
	bc.Save()

	//
	//handlers := api.NewHandlers(&bc)
	//server := api.NewServer(handlers)
	//server.SetupRoutes()
	//server.Run()
}
