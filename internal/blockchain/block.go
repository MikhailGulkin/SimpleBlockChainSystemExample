package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

const (
	MiningDifficulty = 2
	MiningSender     = "THE BLOCKCHAIN"
	MiningReward     = 1.0
)

type Block struct {
	TimeStamp    string         `json:"timeStamp"`
	PrevHash     [32]byte       `json:"prevHash"`
	Transactions []*Transaction `json:"transactions"`
	Nonce        int64          `json:"once"`
}

func NewBlock(nonce int64, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.TimeStamp = time.Now().Format("2006.01.02 15:04:05")
	b.Nonce = nonce
	b.PrevHash = previousHash
	b.Transactions = transactions
	return b
}
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}
func (b *Block) CheckTransactionCompletion(id string) bool {
	for _, t := range b.Transactions {
		if t.Id == id {
			return true
		}
	}
	return false
}
