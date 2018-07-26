// version 1.0 2018-07-26

// ex4.6 squashes adjacent Unicode spaces into a single ASCII space.
package squash

import (
	"unicode"
	"unicode/utf8"
)

func squash(b []byte) []byte {
	w := 0     // write pos
	f := false // space flag
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if !unicode.IsSpace(r) {
			utf8.EncodeRune(b[w:], r)
			w += size
			f = false
		} else {
			if !f {
				size = utf8.EncodeRune(b[w:], int32(' '))
				w += size
			}
			f = true
		}
		i += size
	}

	return b[:w]
}
