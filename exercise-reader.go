package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 'A' 
	}

	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}