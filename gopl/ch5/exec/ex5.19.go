// version 1.0 2018-08-09

// ex5.19 returns a non-zero value using panic and recover
package main

import "fmt"

func weird() (ret string) {
	defer func() {
		recover()
		ret = "hi"
	}()
	panic("omg")
}

func main() {
	fmt.Println(weird())
}
