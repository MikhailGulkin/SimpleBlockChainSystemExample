package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"log"
)

type BlockChain struct {
	Chain               []*Block
	PendingTransactions []*Transaction
	BlockChainAddress   string
}

func NewBlockChain(address string) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.BlockChainAddress = address
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *BlockChain) CreateBlock(once int64, prevHash [32]byte) *Block {
	b := NewBlock(once, prevHash, bc.PendingTransactions)
	bc.PendingTransactions = []*Transaction{}
	bc.Chain = append(bc.Chain, b)
	return b
}
func (bc *BlockChain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *BlockChain) GetBalance(address string) (int64, error) {
	var balance int64
	addressFound := false
	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			if t.FromAddress == address {
				balance -= t.Amount
				addressFound = true
			}
			if t.ToAddress == address {
				balance += t.Amount
				addressFound = true
			}
		}
	}
	if !addressFound {
		return 0, errors.New("address not found")
	}
	return balance, nil
}
func (bc *BlockChain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.PendingTransactions {
		transactions = append(transactions,
			NewTransaction(
				t.FromAddress,
				t.ToAddress,
				t.Amount,
			))
	}
	return transactions
}
func (bc *BlockChain) Load() {
	err := utils.Load(bc, fmt.Sprintf("%s/%s.json", utils.Dir, "block_chain"))
	if err != nil {
		log.Printf("error while loading block chain: %s", err.Error())
	}
	//c := make(chan os.Signal)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//go func() {
	//	<-c
	//	bc.Save()
	//	os.Exit(1)
	//}()
}
func (bc *BlockChain) Save() {
	marshall, err := json.Marshal(&bc)
	if err != nil {
		log.Printf("error while marshalling block chain: %s", err.Error())
	}
	if err := utils.Save(marshall, fmt.Sprintf("%s/%s.json", utils.Dir, "block_chain")); err != nil {
		log.Printf("error while saving block chain: %s", err.Error())
	}
}
