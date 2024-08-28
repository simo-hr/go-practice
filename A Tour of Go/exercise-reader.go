package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read fills the byte slice with ASCII 'A' characters
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
