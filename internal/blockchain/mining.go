package blockchain

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func (bc *BlockChain) Mining() error {
	if len(bc.PendingTransactions) == 0 {
		return errors.New("no transactions")
	}
	_, err := bc.AddTransaction(MiningSender, bc.BlockChainAddress, MiningReward, nil, nil)
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
	guessBlock := Block{time.Time{}, previousHash, transactions, nonce}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	if guessHashStr[:difficulty] == zeros {
		return true
	}
	return false
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

func (bc *BlockChain) StartMining() {
	bc.Mining()

	_ = time.AfterFunc(time.Second*MiningTimerSec, bc.StartMining)
}

func (bc *BlockChain) RegisterNewWallet(blockchainAddress string) bool {

	_, err := bc.AddTransaction(MiningSender, blockchainAddress, 0, nil, nil)

	if err != nil {
		log.Printf("ERROR: %v", err)
		return false
	}
	bc.StartMining()
	return true
}
