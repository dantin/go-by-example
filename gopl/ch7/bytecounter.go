// version 1.0 2018-08-14
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	n := len(p)
	*c += ByteCounter(n) // convert int to ByteCounter
	return n, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprint(&c, "hello, %s", name)
	fmt.Println(c)
}
