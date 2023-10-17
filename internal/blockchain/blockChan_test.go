package blockchain

import (
	"testing"
)

func TestBlockChainCreate(t *testing.T) {
	bc := NewBlockChain(map[string]int64{})
	if len(bc.chain) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.chain))
	}
}
func TestBlockChainTransactionPending(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	bc.PerformTransaction("alice", "bob", 10)
	if len(bc.pendingTransactions) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.pendingTransactions))
	}
	if bc.pendingTransactions[0].fromAddress != "alice" {
		t.Fatalf("expected alice, got: %s", bc.pendingTransactions[0].fromAddress)
	}
	if bc.pendingTransactions[0].toAddress != "bob" {
		t.Fatalf("expected bob, got: %s", bc.pendingTransactions[0].toAddress)
	}
	if bc.wallets["alice"] != 100 {
		t.Fatalf("expected 100, got: %d", bc.wallets["alice"])
	}
	if bc.wallets["bob"] != 100 {
		t.Fatalf("expected 100, got: %d", bc.wallets["bob"])
	}
}
func TestBlockChainProcessPendingTransaction(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	bc.PerformTransaction("alice", "bob", 10)
	bc.ProcessPendingTransaction("alice")
	if len(bc.pendingTransactions) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.pendingTransactions))
	}
	if bc.GetBalance("alice") != 90 {
		t.Fatalf("expected 90, got: %d", bc.GetBalance("alice"))
	}
	if bc.GetBalance("bob") != 110 {
		t.Fatalf("expected 110, got: %d", bc.GetBalance("bob"))
	}
	if len(bc.chain) != 2 {
		t.Fatalf("expected 2, got: %d", len(bc.chain))
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

	bc.ProcessPendingTransaction("alice")
	if len(bc.pendingTransactions) != 1 {
		t.Fatalf("expected 1, got: %d", len(bc.pendingTransactions))
	}
	if bc.GetBalance("alice") != 0 {
		t.Fatalf("expected 0, got: %d", bc.GetBalance("alice"))
	}
	if bc.GetBalance("bob") != 200 {
		t.Fatalf("expected 200, got: %d", bc.GetBalance("bob"))
	}
	if len(bc.chain) != 2 {
		t.Fatalf("expected 2, got: %d", len(bc.chain))
	}

}
func TestBlockChainTransactionCompletion(t *testing.T) {
	bc := NewBlockChain(map[string]int64{
		"alice": 100,
		"bob":   100,
	})
	id := bc.PerformTransaction("alice", "bob", 10)

	errorTxID := bc.PerformTransaction("alice", "bob", 1000) // not enough balance
	bc.ProcessPendingTransaction("alice")

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
	bc.ProcessPendingTransaction("alice")
	bc.PerformTransaction("bob", "alice", 50)
	bc.PerformTransaction("bob", "alice", 50)
	bc.ProcessPendingTransaction("bob")

	if bc.IsValid() != true {
		t.Fatalf("expected true, got: false")
	}

	//corrupt blockchain
	bc.chain[1].transactions[0].amount = 1000
	if bc.IsValid() != false {
		t.Fatalf("blockchain should be invalid")
	}
	// fix blockchain
	bc.chain[1].transactions[0].amount = 10
	if bc.IsValid() != true {
		t.Fatalf("blockchain should be valid")
	}

}
