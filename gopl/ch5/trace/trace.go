// version 1.0 2018-08-08

package main

import (
	"log"
	"time"
)

// bigSlowOperation demo a time consuming function.
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

// trace returns a fuction value that, when called, does the corresponding "on exit" action.
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
