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
		nil,
		nil,
	)
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
	err = bc.Mining()
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
	tx := wallet.NewTransaction(
		MiningSender,
		w.Address,
		w.PrivateKey,
		w.PublicKey,
		10,
	)
	sig := tx.GenerateSignature()
	created, err := bc.AddTransaction(
		MiningSender,
		"toAddress",
		10,
		w.PublicKey,
		sig,
	)
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
	if !created {
		t.Fatalf("expected true, got: %v", created)
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
