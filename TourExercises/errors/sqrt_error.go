package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is an error for negative numbers
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt is an updated version of Sqrt from
// loops/find_sqrt.go that incorporates an error interface
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for z != 0 {
		adjust := (z*z - x) / (2 * z)

		if math.Abs(z-(z-adjust)) > 1e-14 {
			z -= adjust
			fmt.Println(z)
		} else {
			return z, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
