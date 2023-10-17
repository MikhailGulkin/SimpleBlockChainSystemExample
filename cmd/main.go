package main

import (
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/blockchain"
)

func main() {
	var bc blockchain.BlockChain
	defer bc.Save()
	bc.Load()
	fmt.Println(bc.IsValid())

	////defer bc.Save()
	////
	//var bc1 blockchain.BlockChain
	//
	//fmt.Println(bc1.CreateGenesisBlock().TimeStamp.Format("2006.01.02 15:04:05"))
}
