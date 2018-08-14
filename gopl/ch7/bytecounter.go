// version 1.0 2018-08-14
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	length := len(p)
	*c += ByteCounter(length) // convert int to ByteCounter
	return length, nil
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
