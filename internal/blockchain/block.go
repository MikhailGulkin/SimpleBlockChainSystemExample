package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	index        int64
	timeStamp    time.Time
	hash         string
	prevHash     string
	transactions []Transaction
	once         int64
}

func NewBlock(timeStamp time.Time, transactions []Transaction, prevBlock Block) Block {
	return Block{
		index:        prevBlock.index + 1,
		timeStamp:    timeStamp,
		prevHash:     prevBlock.hash,
		transactions: transactions,
		hash:         DefaultHash,
	}
}

func NewGenesisBlock(timeStamp time.Time) Block {
	block := Block{
		index:        0,
		timeStamp:    timeStamp,
		prevHash:     "",
		transactions: make([]Transaction, 0),
		once:         0,
		hash:         DefaultHash,
	}
	block.calculateHash()
	return block
}

func (b *Block) Compare(block Block) bool {
	return b.index == block.index &&
		b.hash == block.hash &&
		b.prevHash == block.prevHash &&
		b.timeStamp == block.timeStamp &&
		TransactionsToString(b.transactions) == TransactionsToString(block.transactions) &&
		b.once == block.once
}
func (b *Block) Mine() {
	for b.IsValid() == false {
		b.once++
		b.calculateHash()
	}
}
func (b *Block) IsValid() bool {
	b.calculateHash()
	return b.hash[0:2] == strings.Repeat("0", Difficulty)
}

func (b *Block) calculateHash() {
	data := fmt.Sprintf("%d%s%s%s%d",
		b.index,
		b.prevHash,
		b.timeStamp.String(),
		TransactionsToString(b.transactions),
		b.once,
	)
	hash := sha256.Sum256([]byte(data))
	b.hash = hex.EncodeToString(hash[:])
}
func (b *Block) IsContainsTxByID(id string) bool {
	for _, tx := range b.transactions {
		if tx.id == id {
			return true
		}
	}
	return false
}
