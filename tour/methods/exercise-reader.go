// version 1.0 2018-07-03
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Add a Reader([]byte) (int, error) method to MyReader
func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 'A'
	}
	return len(bytes), nil

}

func main() {
	reader.Validate(MyReader{})
}
