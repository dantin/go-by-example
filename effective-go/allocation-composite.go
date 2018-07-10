// version 1.0 2018-07-09
package main

import "fmt"

func main() {
	Enone := 0
	Eio := 1
	Einval := 2

	a := [...]string{"no error", "Eio", "invalid argument"}
	s := []string{"no error", "Eio", "invalid argument"}
	//s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(m)
}
