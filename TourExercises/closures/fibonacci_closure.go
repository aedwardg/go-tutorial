package main

import "fmt"

func fibonacci() func() int {
	f2, f1 := 0, 1
	return func() int {
		f := f2

		f2, f1 = f1, f+f1

		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
