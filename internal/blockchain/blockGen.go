package blockchain

import (
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/wallet"
	"io"
	"log"
	"math/rand"
)

const TxNumInBlocks = 100

// GenerateBlocks For generate test data only
func (bc *BlockChain) GenerateBlocks(count int) []*Block {
	log.SetOutput(io.Discard)
	// 2 blocks
	user1 := wallet.NewWallet()

	bc.RegisterNewWallet(user1.Address)

	var ranNum int
	var tx *wallet.Transaction
	var sig *utils.Signature

	for i := 0; i < count; i++ {
		for j := 0; j < TxNumInBlocks; j++ {
			ranNum = rand.Int() % 3
			if ranNum == 0 {
				tx = wallet.NewTransaction(MiningSender, user1.Address, utils.UserTransactionMin, user1.PrivateKey, user1.PublicKey, 5)
				sig = tx.GenerateSignature()
				bc.AddTransaction(MiningSender, user1.Address, 5, utils.UserTransactionMin, user1.PublicKey, sig)
			} else if ranNum == 1 {
				tx = wallet.NewTransaction(MiningSender, user1.Address, utils.UserTransactionAvg, user1.PrivateKey, user1.PublicKey, 55)
				sig = tx.GenerateSignature()
				bc.AddTransaction(MiningSender, user1.Address, 55, utils.UserTransactionAvg, user1.PublicKey, sig)
			} else {
				tx = wallet.NewTransaction(MiningSender, user1.Address, utils.UserTransactionMax, user1.PrivateKey, user1.PublicKey, 555)
				sig = tx.GenerateSignature()
				bc.AddTransaction(MiningSender, user1.Address, 555, utils.UserTransactionMax, user1.PublicKey, sig)
			}
		}
		bc.genMining(bc.BlockChainAddress)
	}
	return bc.Chain[len(bc.Chain)-count:]
}
