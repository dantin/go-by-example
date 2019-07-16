package main

import (
	"flag"
	"fmt"

	"github.com/dantin/go-by-example/gopl/ch2/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
