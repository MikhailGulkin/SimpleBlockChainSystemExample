package blockchain

import (
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/wallet"
	"testing"
)

func TestBlockChain_AddTransaction(t *testing.T) {
	w := wallet.NewWallet()
	bcWallet := wallet.NewWallet()
	bc := NewBlockChain(bcWallet.Address)
	_, err := bc.AddTransaction(
		MiningSender,
		w.Address,
		10,
		wallet.GetTransactionType(10),
		nil,
		nil,
	)
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
	err = bc.Mining(bc.BlockChainAddress)
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
	tx := wallet.NewTransaction(
		MiningSender,
		w.Address,
		wallet.GetTransactionType(10),
		w.PrivateKey,
		w.PublicKey,
		10,
	)
	sig := tx.GenerateSignature()
	id, err := bc.AddTransaction(
		MiningSender,
		"toAddress",
		10,
		wallet.GetTransactionType(10),
		w.PublicKey,
		sig,
	)
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
	if id == "" {
		t.Fatalf("expected not empty id, got: %v", id)
	}

	t.Run("get balance", func(t *testing.T) {
		balance, err := bc.GetBalance(w.Address)
		if err != nil {
			t.Fatalf("expected nil, got: %v", err)
		}
		if balance != 10 {
			t.Fatalf("expected 10, got: %v", balance)
		}
	})
}

func TestBlockChain_GenerateBlocks(t *testing.T) {
	w := wallet.NewWallet()
	bc := NewBlockChain(w.Address)
	bc.RegisterNewWallet(wallet.NewWallet().Address)
	bc.RegisterNewWallet(wallet.NewWallet().Address)
	bc.RegisterNewWallet(wallet.NewWallet().Address)
	bc.RegisterNewWallet(wallet.NewWallet().Address)

	prevLen := len(bc.Chain)
	blockGenCount := 10
	blocks := bc.GenerateBlocks(blockGenCount)

	if len(blocks) != blockGenCount {
		t.Fatalf("!=")
	}
	if len(bc.Chain) != blockGenCount+1+prevLen {
		t.Fatalf("")
	}
	for _, block := range blocks {
		if len(block.Transactions) != TxNumInBlocks {
			t.Fatalf("")
		}
	}
}
