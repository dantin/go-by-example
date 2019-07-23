package bank

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance.
	balance int
)

// Deposit save an amount to deposit.
func Deposit(amount int) {
	// acquire token
	sema <- struct{}{}
	balance = balance + amount
	// release token
	<-sema
}

// Balance return the balance amount of an account.
func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
