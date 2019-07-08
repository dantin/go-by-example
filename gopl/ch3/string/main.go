// version 1.0 2018-07-23

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// The build-in `len` function returns the number of bytes (not rune) in a string.
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	// Hwo to process each character using UTF-8 decoder.
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// The range loop performs UTF-8 decoding implicityly.
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// A `[]rune` conversion apply to a UTF-8 encoded string return the sequence of Unicode code point.
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))

	// The special Unicode replacement character, '\uFFFD', BAD UTF-8 rune.
	fmt.Println(string(1234567))
}
