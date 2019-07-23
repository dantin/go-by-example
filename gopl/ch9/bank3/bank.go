package bank

import "sync"

var (
	// gurads balance
	mu      sync.Mutex
	balance int
)

// Deposit save an amount to deposit.
func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

// Balance return the balance amount of an account.
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
