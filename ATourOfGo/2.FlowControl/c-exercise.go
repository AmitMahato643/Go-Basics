package main

import "fmt"

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 20; i++ {
		fmt.Printf("Z = %v\n", z)
		z -= (z*z - x) / 2 * z
	}
	return z
}

func main() {
	Sqrt(4)
}
