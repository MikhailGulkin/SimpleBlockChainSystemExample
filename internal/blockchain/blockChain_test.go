package blockchain

import (
	"github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/utils"
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
		utils.UserTransaction,
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
		utils.UserTransaction,
		w.PrivateKey,
		w.PublicKey,
		10,
	)
	sig := tx.GenerateSignature()
	id, err := bc.AddTransaction(
		MiningSender,
		"toAddress",
		10,
		utils.UserTransaction,
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
