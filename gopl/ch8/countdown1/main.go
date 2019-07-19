package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("launch")
}

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countown := 10; countown > 0; countown-- {
		fmt.Println(countown)
		<-tick
	}
	launch()
}
