package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BlockChain struct {
	Chain               []Block          `json:"chain"`
	PendingTransactions []Transaction    `json:"pendingTransactions"`
	Wallets             map[string]int64 `json:"wallets"`
	Reward              int64            `json:"reward"`
}

func NewBlockChain(wallets map[string]int64) BlockChain {
	bc := BlockChain{
		Wallets:             wallets,
		Reward:              1,
		PendingTransactions: make([]Transaction, 0),
	}
	bc.createGenesisBlock()
	return bc
}

func (bc *BlockChain) createGenesisBlock() Block {
	genesisBlock := NewGenesisBlock(time.Now())
	genesisBlock.Mine()
	bc.Chain = append(bc.Chain, genesisBlock)
	return bc.Chain[0]
}
func (bc *BlockChain) getLatestBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *BlockChain) AddBlock(block Block) {
	block.Mine()
	bc.Chain = append(bc.Chain, block)
}
func (bc *BlockChain) createTransaction(transaction Transaction) {
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)
}

func (bc *BlockChain) ProcessPendingTransaction(mineAddress string) error {
	if len(bc.PendingTransactions) == 0 {
		return errors.New("no pending transactions")
	}
	if mineAddress == "" {
		return errors.New("mine address is empty")
	}
	if _, ok := bc.Wallets[mineAddress]; !ok {
		return fmt.Errorf("mine address %s is not found", mineAddress)
	}
	processed := bc.ProcessTransactions(bc.PendingTransactions)
	block := NewBlock(time.Now(), processed, bc.getLatestBlock())
	bc.AddBlock(block)
	bc.PendingTransactions = []Transaction{NewTransaction("", mineAddress, bc.Reward)}
	return nil
}
func (bc *BlockChain) GetBalance(address string) int64 {
	return bc.Wallets[address]
}
func (bc *BlockChain) GetWallets() map[string]int64 {
	return bc.Wallets
}
func (bc *BlockChain) ProcessTransactions(transaction []Transaction) []Transaction {
	processedTransactions := make([]Transaction, 0)
	for _, tx := range transaction {
		if tx.FromAddress == "" {
			bc.Wallets[tx.ToAddress] += tx.Amount
			processedTransactions = append(processedTransactions, tx)
		} else if isBalanceSufficient(bc.Wallets[tx.FromAddress], tx.Amount) {
			bc.Wallets[tx.FromAddress] -= tx.Amount
			bc.Wallets[tx.ToAddress] += tx.Amount
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
	return transaction.Id
}
func (bc *BlockChain) IsValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		if bc.Chain[i].IsValid() == false {
			return false
		}
		if bc.Chain[i].PrevHash != bc.Chain[i-1].Hash {
			return false
		}
	}
	return true
}
func (bc *BlockChain) Save() {
	marshall, err := json.Marshal(&bc)
	if err != nil {
		log.Printf("error while marshalling block chain: %s", err.Error())
	}
	if err := Save(marshall, fmt.Sprintf("%s/%s.json", Dir, "block_chain")); err != nil {
		log.Printf("error while saving block chain: %s", err.Error())
	}
}
func (bc *BlockChain) Load() {
	err := Load(bc, fmt.Sprintf("%s/%s.json", Dir, "block_chain"))
	if err != nil {
		log.Printf("error while loading block chain: %s", err.Error())
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		bc.Save()
		os.Exit(1)
	}()
}
func (bc *BlockChain) CheckTransactionCompletion(id string) bool {
	for _, block := range bc.Chain {
		if block.IsContainsTxByID(id) {
			return true
		}
	}
	return false
}
