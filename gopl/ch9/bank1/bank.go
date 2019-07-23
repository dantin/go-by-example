package bank

// Package bank provides a concurrency-safe bank with one account.

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

// Deposit save an amount to deposit.
func Deposit(amount int) {
	deposits <- amount
}

// Balance return the balance amount of an account.
func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine.
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	// start the monitor goroutine.
	go teller()
}
