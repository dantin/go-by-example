// version 1.0 2018-07-10
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Println(i, v)
	}
}
