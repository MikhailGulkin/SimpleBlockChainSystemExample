package blockchain

import (
	"errors"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"log"
	"strings"
)

func (bc *BlockChain) Mining(address string) error {
	if len(bc.PendingTransactions) == 0 {
		return errors.New("no transactions")
	}
	if address == "" {
		return errors.New("address is empty")
	}
	if _, ok := bc.GetWallets()[address]; !ok && address != bc.BlockChainAddress {
		return errors.New("invalid address")
	}
	_, err := bc.AddTransaction(MiningSender, address, MiningReward, utils.Mining, nil, nil)
	if err != nil {
		return err
	}
	nonce := bc.ProofOfWork()
	prevHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, prevHash)
	return nil
}

func (bc *BlockChain) ProofOfWork() int64 {
	txs := bc.CopyTransactionPool()
	prevHash := bc.LastBlock().Hash()
	nonce := int64(0)
	for !bc.ValidProof(nonce, prevHash, txs, MiningDifficulty) {
		nonce++
	}
	return nonce
}
func (bc *BlockChain) ValidProof(nonce int64, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{"", previousHash, transactions, nonce}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}
func (bc *BlockChain) ValidChain(chain []*Block) bool {
	preBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		block := chain[currentIndex]
		if block.PrevHash != preBlock.Hash() {
			return false
		}
		if !bc.ValidProof(block.Nonce, block.PrevHash, block.Transactions, MiningDifficulty) {
			return false
		}
		preBlock = block
		currentIndex++
	}
	return true

}

func (bc *BlockChain) RegisterNewWallet(blockchainAddress string) bool {
	_, err := bc.AddTransaction(MiningSender, blockchainAddress, 0, utils.WalletCreate, nil, nil)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return false
	}
	err = bc.Mining(bc.BlockChainAddress)

	return err == nil
}
