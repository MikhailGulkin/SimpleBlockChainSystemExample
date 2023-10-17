package blockchain

import (
	"log"
	"time"
)

type BlockChain struct {
	chain               []Block
	pendingTransactions []Transaction
	wallets             map[string]int64
	reward              int64
}

func NewBlockChain(wallets map[string]int64) BlockChain {
	bc := BlockChain{
		wallets:             wallets,
		reward:              1,
		pendingTransactions: make([]Transaction, 0),
	}
	bc.CreateGenesisBlock()
	return bc
}

func (bc *BlockChain) CreateGenesisBlock() Block {
	genesisBlock := NewGenesisBlock(time.Now())
	genesisBlock.Mine()
	bc.chain = append(bc.chain, genesisBlock)
	return bc.chain[0]
}
func (bc *BlockChain) getLatestBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) AddBlock(block Block) {
	block.Mine()
	bc.chain = append(bc.chain, block)
}
func (bc *BlockChain) createTransaction(transaction Transaction) {
	bc.pendingTransactions = append(bc.pendingTransactions, transaction)
}

func (bc *BlockChain) ProcessPendingTransaction(mineAddress string) {
	processed := bc.ProcessTransactions(bc.pendingTransactions)
	block := NewBlock(time.Now(), processed, bc.getLatestBlock())
	bc.AddBlock(block)
	bc.pendingTransactions = []Transaction{NewTransaction("", mineAddress, bc.reward)}
}
func (bc *BlockChain) GetBalance(address string) int64 {
	return bc.wallets[address]
}
func (bc *BlockChain) ProcessTransactions(transaction []Transaction) []Transaction {
	processedTransactions := make([]Transaction, 0)
	for _, tx := range transaction {
		if tx.fromAddress == "" {
			bc.wallets[tx.toAddress] += tx.amount
			processedTransactions = append(processedTransactions, tx)
		} else if isBalanceSufficient(bc.wallets[tx.fromAddress], tx.amount) {
			bc.wallets[tx.fromAddress] -= tx.amount
			bc.wallets[tx.toAddress] += tx.amount
			processedTransactions = append(processedTransactions, tx)
		} else {
			log.Printf("balance is not sufficient for transaction: %s", tx.ToString())
		}
	}
	return processedTransactions
}

func (bc *BlockChain) PerformTransaction(sender, receiver string, amount int64) string {
	transaction := NewTransaction(sender, receiver, amount)
	bc.createTransaction(transaction)
	return transaction.id
}
func (bc *BlockChain) IsValid() bool {
	for i := 1; i < len(bc.chain); i++ {
		if bc.chain[i].IsValid() == false {
			return false
		}
		if bc.chain[i].prevHash != bc.chain[i-1].hash {
			return false
		}
	}
	return true
}

func (bc *BlockChain) CheckTransactionCompletion(id string) bool {
	for _, block := range bc.chain {
		if block.IsContainsTxByID(id) {
			return true
		}
	}
	return false
}
