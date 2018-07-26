// version 1.0 2018-07-25

// ex4.7 reverse an UTF-8 string in-place.
package main

import "unicode/utf8"

func rev(b []byte) {
	size := len(b)
	for i := 0; i < size/2; i++ {
		b[i], b[size-i-1] = b[size-i-1], b[i]
	}
}

func revUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}
