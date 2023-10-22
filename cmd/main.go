package main

import (
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/wallet"
)

func main() {
	var bc blockchain.BlockChain
	bc.Load()
	var wallets wallet.Wallets
	wallets.Load()

	handlers := api.NewHandlers(&wallets, &bc)
	server := api.NewServer(handlers)
	server.SetupRoutes()
	server.Run()
	//
	//wallets := wallet.NewWallets()
	//w := wallets.NewWallet()
	//user1 := wallets.NewWallet()
	//user2 := wallets.NewWallet()
	//wallets.Save()
	//bc := blockchain.NewBlockChain(w.Address)
	//bc.RegisterNewWallet(user1.Address)
	//bc.RegisterNewWallet(user2.Address)
	//bc.AddTransaction(
	//	blockchain.MiningSender,
	//	user1.Address,
	//	100,
	//	utils.UserTransaction,
	//	nil,
	//	nil,
	//)
	//bc.AddTransaction(
	//	blockchain.MiningSender,
	//	user2.Address,
	//	100,
	//	utils.UserTransaction,
	//	nil,
	//	nil,
	//)
	//bc.Mining(user1.Address)
	//fmt.Println(bc.GetWallets())
	//fmt.Println(bc.ValidChain(bc.Chain))
	//bc.Save()
	//
	//var wallets wallet.Wallets
	//var bc blockchain.BlockChain
	//wallets.Load()
	//user1 := wallets.Wallets[1]
	//user2 := wallets.Wallets[2]
	////wallets.Save()
	//bc.Load()
	//fmt.Println(bc.GetWallets(), user1.Address)
	//tx := wallet.NewTransaction(
	//	user1.Address,
	//	user2.Address,
	//	utils.UserTransaction,
	//	user1.PrivateKey,
	//	user1.PublicKey,
	//	91,
	//)
	//sig := tx.GenerateSignature()
	//_, err := bc.AddTransaction(
	//	user1.Address,
	//	user2.Address,
	//	91,
	//	utils.UserTransaction,
	//	user1.PublicKey,
	//	sig,
	//)
	//if err != nil {
	//	panic(err)
	//}
	//bc.Mining(user1.Address)
	//fmt.Println(bc.GetWallets(), user1.Address)
	//fmt.Println(bc.ValidChain(bc.Chain))
	//bc.Save()
}
