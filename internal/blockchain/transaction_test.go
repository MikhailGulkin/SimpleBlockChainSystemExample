package blockchain

import (
	"testing"
)

func TestGenerateTransactionId(t *testing.T) {
	randId := GenerateTransactionId()

	if len(randId) != 24 {
		t.Fatalf("expected 24, got: %d", len(randId))
	}
}
