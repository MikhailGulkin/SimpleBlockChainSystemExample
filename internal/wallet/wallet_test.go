package wallet

import (
	"testing"
)

func TestNewWallet(t *testing.T) {
	wallet := NewWallet()
	if wallet.Address == "" {
		t.Fatalf("expected not empty, got: %s", wallet.Address)
	}
	if wallet.PrivateKey == nil {
		t.Fatalf("expected not nil, got: %v", wallet.PrivateKey)
	}
	if wallet.PublicKey == nil {
		t.Fatalf("expected not nil, got: %v", wallet.PublicKey)
	}
}
