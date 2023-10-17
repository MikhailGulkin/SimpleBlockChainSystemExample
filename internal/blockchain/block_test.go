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
	if block.hash != exceptedHash {
		t.Fatalf("expected hash: %s, got: %s", exceptedHash, block.hash)
	}
}
func TestBlockMine(t *testing.T) {
	block := NewBlock(time.Now(), []Transaction{}, Block{})
	if block.once != 0 {
		t.Fatalf("expected once: 0, got: %d", block.once)
	}
	block.Mine()
	if block.once == 0 {
		t.Fatalf("expected once: more then 0, got: %d", block.once)
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

	if genesisBlock.index != 0 {
		t.Fatalf("expected index: 0, got: %d", genesisBlock.index)
	}
	if genesisBlock.prevHash != "" {
		t.Fatalf("expected prevHash: \"\", got: %s", genesisBlock.prevHash)
	}
	if genesisBlock.once != 0 {
		t.Fatalf("expected once: 0, got: %d", genesisBlock.once)
	}
	if genesisBlock.hash == "" {
		t.Fatalf("expected hash: not empty, got: %s", genesisBlock.hash)
	}
	if genesisBlock.hash != exceptedHash {
		t.Fatalf("expected hash: %s, got: %s", exceptedHash, genesisBlock.hash)
	}
	if genesisBlock.timeStamp != timeStamp {
		t.Fatalf("expected timeStamp: %s, got: %s", timeStamp, genesisBlock.timeStamp)
	}
	if len(genesisBlock.transactions) != 0 {
		t.Fatalf("expected transactions: 0, got: %d", len(genesisBlock.transactions))
	}
}
