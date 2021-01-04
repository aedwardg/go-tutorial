package main

import "golang.org/x/tour/pic"

// Pic returns a slice of length dy, each element of which
// is a slice of dx 8-bit unsigned integers.
func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	for y := range img {
		img[y] = make([]uint8, dx)

		for x := range img[y] {
			img[y][x] = uint8(x ^ y)
		}
	}
	return img
}

func main() {
	// Shows picture defined by function when executed on Go Playground
	pic.Show(Pic)
}
