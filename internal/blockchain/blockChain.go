package blockchain

import (
	"encoding/json"
	"errors"
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type BlockChain struct {
	Chain               []*Block       `json:"chain"`
	PendingTransactions []*Transaction `json:"pendingTransactions"`
	BlockChainAddress   string         `json:"blockChainAddress"`
}

func NewBlockChain(address string) *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.BlockChainAddress = address
	bc.CreateBlock(0, b.Hash())
	return bc
}
func (bc *BlockChain) GetChain() []*Block {
	return bc.Chain
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
func (bc *BlockChain) GetNBlocks(n int) []*Block {
	if n > len(bc.Chain) {
		return bc.Chain
	}
	return bc.Chain[len(bc.Chain)-n:]
}
func (bc *BlockChain) GetFixedBalance(address string) (int64, error) {
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
	for _, t := range bc.PendingTransactions {
		if t.FromAddress == address {
			balance -= t.Amount
			addressFound = true
		}
		if t.ToAddress == address {
			balance += t.Amount
			addressFound = true
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
		tx := NewTransaction(
			t.FromAddress,
			t.ToAddress,
			t.TransactionType,
			t.Amount,
		)
		tx.Id = t.Id
		tx.Time = t.Time
		transactions = append(transactions, tx)
	}
	return transactions
}
func (bc *BlockChain) CheckTransactionCompletion(id string) bool {
	for _, b := range bc.Chain {
		if b.CheckTransactionCompletion(id) {
			return true
		}
	}
	return false

}
func (bc *BlockChain) GetWallets() map[string]int64 {
	set := make(map[string]int64)
	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			if t.TransactionType == utils.WalletCreate {
				set[t.ToAddress] = 0
			}
		}
	}
	for key := range set {
		set[key], _ = bc.GetFixedBalance(key)
	}
	return set
}
func (bc *BlockChain) GetWalletTransactions(address string) []*Transaction {
	var transactions []*Transaction
	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			if t.FromAddress == address || t.ToAddress == address {
				transactions = append(transactions, t)
			}
		}
	}
	return transactions
}

func (bc *BlockChain) Load() {
	err := utils.Load(bc, "block_chain")
	if err != nil {
		log.Printf("error while loading block chain: %s", err.Error())
	}
}
func (bc *BlockChain) SignalSave() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		bc.Save()
		os.Exit(1)
	}()
}
func (bc *BlockChain) Save() {
	marshall, err := json.Marshal(bc)
	if err != nil {
		log.Printf("error while marshalling block chain: %s", err.Error())
	}
	if err := utils.Save(marshall, "block_chain"); err != nil {
		log.Printf("error while saving block chain: %s", err.Error())
	}
}
