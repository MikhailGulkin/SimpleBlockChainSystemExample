package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index        int64         `json:"index"`
	TimeStamp    time.Time     `json:"timeStamp"`
	Hash         string        `json:"hash"`
	PrevHash     string        `json:"prevHash"`
	Transactions []Transaction `json:"transactions"`
	Once         int64         `json:"once"`
}

func NewBlock(timeStamp time.Time, transactions []Transaction, prevBlock Block) Block {
	return Block{
		Index:        prevBlock.Index + 1,
		TimeStamp:    timeStamp,
		PrevHash:     prevBlock.Hash,
		Transactions: transactions,
		Hash:         DefaultHash,
	}
}

func NewGenesisBlock(timeStamp time.Time) Block {
	block := Block{
		Index:        0,
		TimeStamp:    timeStamp,
		PrevHash:     "",
		Transactions: make([]Transaction, 0),
		Once:         0,
		Hash:         DefaultHash,
	}
	block.calculateHash()
	return block
}

func (b *Block) Compare(block Block) bool {
	return b.Index == block.Index &&
		b.Hash == block.Hash &&
		b.PrevHash == block.PrevHash &&
		b.TimeStamp == block.TimeStamp &&
		TransactionsToString(b.Transactions) == TransactionsToString(block.Transactions) &&
		b.Once == block.Once
}
func (b *Block) Mine() {
	for b.IsValid() == false {
		b.Once++
	}
}
func (b *Block) IsValid() bool {
	b.calculateHash()
	return b.Hash[0:Difficulty] == strings.Repeat("0", Difficulty)
}

func (b *Block) calculateHash() {

	data := fmt.Sprintf("%d%s%s%s%d",
		b.Index,
		b.PrevHash,
		b.TimeStamp.Format("2006.01.02 15:04:05"),
		TransactionsToString(b.Transactions),
		b.Once,
	)
	hash := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hash[:])
}
func (b *Block) IsContainsTxByID(id string) bool {
	for _, tx := range b.Transactions {
		if tx.Id == id {
			return true
		}
	}
	return false
}
