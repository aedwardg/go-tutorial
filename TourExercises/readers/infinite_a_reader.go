package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Exercise: Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

func (reader MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
