package blockchain

func isBalanceSufficient(senderBalance int64, amount int64) bool {
	return senderBalance >= amount
}
