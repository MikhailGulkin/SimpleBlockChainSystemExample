package blockchain

import (
	"testing"
)

func TestBlockChainCreate(t *testing.T) {
	bc := NewBlockChain(map[string]int64{})
	if len(bc.Chain) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.Chain))
	}
}
func TestBlockChainTransactionPending(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	bc.PerformTransaction("alice", "bob", 10)
	if len(bc.PendingTransactions) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.PendingTransactions))
	}
	if bc.PendingTransactions[0].FromAddress != "alice" {
		t.Fatalf("expected alice, got: %s", bc.PendingTransactions[0].FromAddress)
	}
	if bc.PendingTransactions[0].ToAddress != "bob" {
		t.Fatalf("expected bob, got: %s", bc.PendingTransactions[0].ToAddress)
	}
	if bc.Wallets["alice"] != 100 {
		t.Fatalf("expected 100, got: %d", bc.Wallets["alice"])
	}
	if bc.Wallets["bob"] != 100 {
		t.Fatalf("expected 100, got: %d", bc.Wallets["bob"])
	}
}
func TestBlockChainProcessPendingTransaction(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	bc.PerformTransaction("alice", "bob", 10)
	err := bc.ProcessPendingTransaction("alice")
	if err != nil {
		t.Fatalf("expected nil, got: %s", err.Error())
	}
	if len(bc.PendingTransactions) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.PendingTransactions))
	}
	if bc.GetBalance("alice") != 90 {
		t.Fatalf("expected 90, got: %d", bc.GetBalance("alice"))
	}
	if bc.GetBalance("bob") != 110 {
		t.Fatalf("expected 110, got: %d", bc.GetBalance("bob"))
	}
	if len(bc.Chain) != 2 {
		t.Fatalf("expected 2, got: %d", len(bc.Chain))
	}
}
func TestBlockChainProcessPendingMultiplyTransaction(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	bc.PerformTransaction("alice", "bob", 10)
	bc.PerformTransaction("alice", "bob", 90)

	bc.PerformTransaction("alice", "bob", 1000) // not enough balance

	err := bc.ProcessPendingTransaction("alice")
	if err != nil {
		t.Fatalf("expected nil, got: %s", err.Error())
	}
	if len(bc.PendingTransactions) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.PendingTransactions))
	}
	if bc.GetBalance("alice") != 0 {
		t.Fatalf("expected 0, got: %d", bc.GetBalance("alice"))
	}
	if bc.GetBalance("bob") != 200 {
		t.Fatalf("expected 200, got: %d", bc.GetBalance("bob"))
	}
	if len(bc.Chain) != 2 {
		t.Fatalf("expected 2, got: %d", len(bc.Chain))
	}

}
func TestBlockChainTransactionCompletion(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	id := bc.PerformTransaction("alice", "bob", 10)

	errorTxID := bc.PerformTransaction("alice", "bob", 1000) // not enough balance
	err := bc.ProcessPendingTransaction("alice")
	if err != nil {
		t.Fatalf("expected nil, got: %s", err.Error())
	}

	if bc.CheckTransactionCompletion(id) != true {
		t.Fatalf("expected true, got: false")
	}
	if bc.CheckTransactionCompletion(errorTxID) != false {
		t.Fatalf("expected false, got: true")
	}
}
func TestBlockChainValidation(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	bc.PerformTransaction("alice", "bob", 10)
	bc.PerformTransaction("alice", "bob", 90)
	err := bc.ProcessPendingTransaction("alice")
	if err != nil {
		t.Fatalf("expected nil, got: %s", err.Error())
	}
	bc.PerformTransaction("bob", "alice", 50)
	bc.PerformTransaction("bob", "alice", 50)
	err = bc.ProcessPendingTransaction("bob")
	if err != nil {
		t.Fatalf("expected nil, got: %s", err.Error())
	}

	if bc.IsValid() != true {
		t.Fatalf("expected true, got: false")
	}

	//corrupt blockchain
	bc.Chain[1].Transactions[0].Amount = 1000
	if bc.IsValid() != false {
		t.Fatalf("blockchain should be invalid")
	}
	// fix blockchain
	bc.Chain[1].Transactions[0].Amount = 10
	if bc.IsValid() != true {
		t.Fatalf("blockchain should be valid")
	}

}

func TestBlockChainProcessingEror(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	bc.PerformTransaction("alice", "bob", 10)
	bc.PerformTransaction("alice", "bob", 90)

	if err := bc.ProcessPendingTransaction("some"); err == nil {
		t.Fatalf("except error, got: nil")
	}
	if err := bc.ProcessPendingTransaction("alice"); err != nil {
		t.Fatalf("except nil, got: %s", err.Error())
	}
	bc = NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	if err := bc.ProcessPendingTransaction("alice"); err == nil {
		t.Fatalf("except error, got: nil")
	}

}
