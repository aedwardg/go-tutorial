package main

import (
	"fmt"
	"math"
)

// Sqrt tries to find the square root of a number
// using Newton's method.
func Sqrt(x float64) float64 {
	z := 1.0
	for z != 0 {
		adjust := (z*z - x) / (2 * z)

		if math.Abs(z-(z-adjust)) > 1e-14 {
			z -= adjust
			fmt.Println(z)
		} else {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(225), math.Sqrt(225))
}
