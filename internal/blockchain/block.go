package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

const (
	MiningDifficulty = 3
	MiningSender     = "THE BLOCKCHAIN"
	MiningReward     = 1.0
	MINING_TIMER_SEC = 20
)

type Block struct {
	TimeStamp    time.Time      `json:"timeStamp"`
	PrevHash     [32]byte       `json:"prevHash"`
	Transactions []*Transaction `json:"transactions"`
	Nonce        int64          `json:"once"`
}

func NewBlock(nonce int64, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.TimeStamp = time.Now()
	b.Nonce = nonce
	b.PrevHash = previousHash
	b.Transactions = transactions
	return b
}
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}
