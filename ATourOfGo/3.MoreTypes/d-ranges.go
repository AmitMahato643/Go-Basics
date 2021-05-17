package main

import "fmt"

func main() {
	// 16. Range - the range form of a for loop iterate over a slice or a map
	// when ranging over a slice or map, return two values, an index and the value resp for each iteration.

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// index or value can be omitted using _ in its place
	for i, v := range slice {
		fmt.Printf("Iteration %v: %v * %v = %v\n", i+1, v, v, v*v)
	}

	// 17. Omit values of range that are not needed
	// for _, v := range slice ; omits index
	// for i, _ := range slice or for i := range slice ; omits values

	for _, v := range slice {
		fmt.Printf("%v, ", v)
	}
	fmt.Println()

}
