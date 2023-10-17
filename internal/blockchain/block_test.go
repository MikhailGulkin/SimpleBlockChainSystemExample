package blockchain

import (
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	transaction := Transaction{}
	timeStamp, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", "2021-01-01 00:00:00 +0000 UTC")
	if err != nil {
		t.Fatal(err)
	}
	block := NewBlock(timeStamp, []Transaction{transaction}, Block{})

	exceptedHash := "d7a2dfc97d83691cbaa8c1a88f8191f28add4847bd25fb1b361d950a934d2807"
	block.calculateHash()
	if block.Hash != exceptedHash {
		t.Fatalf("expected hash: %s, got: %s", exceptedHash, block.Hash)
	}
}
func TestBlockMine(t *testing.T) {
	block := NewBlock(time.Now(), []Transaction{}, Block{})
	if block.Once != 0 {
		t.Fatalf("expected once: 0, got: %d", block.Once)
	}
	block.Mine()
	if block.Once == 0 {
		t.Fatalf("expected once: more then 0, got: %d", block.Once)
	}
}
func TestBlockIsValid(t *testing.T) {
	block := NewBlock(time.Now(), []Transaction{}, Block{})
	if block.IsValid() != false {
		t.Fatalf("expected true, got: false")
	}
	block.Mine()
	if block.IsValid() != true {
		t.Fatalf("expected false, got: true")
	}

}

func TestBlockCompare(t *testing.T) {
	block := NewBlock(time.Now(), []Transaction{}, Block{})
	if block.Compare(block) == false {
		t.Fatalf("expected true, got: false")
	}
}
func TestNewGenesisBlock(t *testing.T) {
	timeStamp, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", "2021-01-01 00:00:00 +0000 UTC")
	if err != nil {
		t.Fatal(err)
	}
	genesisBlock := NewGenesisBlock(timeStamp)
	exceptedHash := "d3b5d19ceffdcf5847599fdc55ea39af35817ea0bac8f1346ef20b9dc2b62c00"

	if genesisBlock.Index != 0 {
		t.Fatalf("expected index: 0, got: %d", genesisBlock.Index)
	}
	if genesisBlock.PrevHash != "" {
		t.Fatalf("expected prevHash: \"\", got: %s", genesisBlock.PrevHash)
	}
	if genesisBlock.Once != 0 {
		t.Fatalf("expected once: 0, got: %d", genesisBlock.Once)
	}
	if genesisBlock.Hash == "" {
		t.Fatalf("expected hash: not empty, got: %s", genesisBlock.Hash)
	}
	if genesisBlock.Hash != exceptedHash {
		t.Fatalf("expected hash: %s, got: %s", exceptedHash, genesisBlock.Hash)
	}
	if genesisBlock.TimeStamp != timeStamp {
		t.Fatalf("expected timeStamp: %s, got: %s", timeStamp, genesisBlock.TimeStamp)
	}
	if len(genesisBlock.Transactions) != 0 {
		t.Fatalf("expected transactions: 0, got: %d", len(genesisBlock.Transactions))
	}
}
