// version 1.0 2018-07-06
package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}
