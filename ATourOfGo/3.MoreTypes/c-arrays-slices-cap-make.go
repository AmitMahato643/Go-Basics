package main

import "fmt"

func main() {
	// 6. Arrays : type [n]T is an array of n values of type T
	// the length of array is the part of its type and hence an array is not resizable
	// First way of creating an array, initialization after declaration - one by one
	var firstTwoPrimes [2]int
	firstTwoPrimes[0] = 1
	firstTwoPrimes[1] = 2

	fmt.Println("First Two Primes : ", firstTwoPrimes)

	// second method of creating an array, initilization declaration
	names := [4]string{"ABC", "XYZ", "PQR", "MNO"}
	fmt.Println("Names : ", names)

	// array of struct type, initialization at declaration
	Vertices := [3]struct {
		x, y int
	}{
		{1, 2},
		{2, 3},
		{3, 4},
	}

	fmt.Println("Vertices of a triangle : ", Vertices)

	type Car struct {
		name  string // zero value of ""
		wheel int    // zero value of 0
	}
	var Cars [3]Car
	Cars[0] = Car{"Tesla", 4}
	// remaininig two will have zero values

	fmt.Println("Car name and their wheel count : ", Cars)

	// 7. Slices
	// Slices are of dynamic size unlike array
	// type []T is a slice with elements of type T
	// a slice is formed by specifying two indices, a low and high bound, separated by colon
	primes := [8]int{1, 2, 3, 5, 7, 11, 13, 17} // array literal of prime numbers
	firstFivePrimes := primes[0:5]              // a slice of prime numbers from index 0 to 5 (excludes element at 5th index)
	primesStartsFromThree := primes[2:]

	fmt.Println()
	fmt.Println("Prime array : ", primes)
	fmt.Println("Slice of first five elements of prime array : ", firstFivePrimes)
	fmt.Println("Slice of elements starting from 3rd element in prime array : ", primesStartsFromThree)

	// 8. Slices are like reference to array
	// in above two slices elements from index 2 to 4 are common
	// and modifying any element refering to primes at that index will have impact on all places
	firstFivePrimes[3] = 19

	fmt.Println()
	fmt.Println("Prime array : ", primes)
	fmt.Println("Slice of first five elements of prime array : ", firstFivePrimes)
	fmt.Println("Slice of elements starting from 3rd element in prime array : ", primesStartsFromThree)

	// 9. Slice and array literals
	randomSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // creates an array first and references it to the slice
	fmt.Println()
	fmt.Println("Random slice literal : ", randomSlice)

	// 10. slice defaults
	// when slicing the higher and lower bound can be omitted to use their defaults
	// From above randomSlice, examples are shown below
	fmt.Println()
	fmt.Println("randomSlice[0:9] 	= ", randomSlice[0:9])
	fmt.Println("randomSlice[:9] 	= ", randomSlice[:9])
	fmt.Println("randomSlice[0:] 	= ", randomSlice[0:])
	fmt.Println("randomSlice[:] 		= ", randomSlice[:])

	// !important
	// 11. Slice length and capacity
	// length of slice can be found from len(slice) - length of slice is the number of elements in the slice
	// capacity = cap(slice) - capacity of slice is the number of elements in the underlying array, counting from "first element in the slice"
	// a slice can be extended by re-slicing it upto maximum capacity - capacity depends on where the first element of slice indexes to the underlying array
	CarBrands := [8]string{"Acura", "Alfa", "Audi", "BMW", "Bentley", "Buick", "Cadillac", "Chevrolet"}

	fmt.Println()
	fmt.Println(CarBrands)
	// slice to give zero length, sarts and ends at index 0, hence max capacity
	slice := CarBrands[:0]
	printStringSlice(slice)

	// extend to get first 5 car brands, still starts at 0 index hence max capacity
	slice = slice[:5]
	printStringSlice(slice)

	// drops first two car brands, starts at 2 hence capacity = max capacity - 2
	slice = slice[2:] // still upper bound follows previous slice
	printStringSlice(slice)

	// extend to get all car brands, but only from 3rd element i.e. from index 2
	slice = slice[:cap(slice)]
	printStringSlice(slice)

	// 12. A Nil Slice - The zero value of a slice is Nil and has len and cap both 0
	var s []int
	fmt.Println()
	fmt.Println("A Nil slice : ", s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	// 13. Creating a slice with make
	// in-built function make to create dynamically-sized array
	// the make function allocates a zero values array and returns a slice referencing to that array
	makeSlice := make([]int, 5)           // len = 5, cap = 5
	makeSliceCapTen := make([]int, 2, 10) // len = 2, cap = 10

	fmt.Println()
	printIntSlice(makeSlice)
	printIntSlice(makeSliceCapTen)

	// 14. slices of slice
	slices := [][]int{
		{2, 4},
		{5, 6},
		{4, 9},
	} // like a 2x2 matrix
	makeSlices := make([][]int, 5, 7)
	makeSlices[0] = make([]int, 2)

	fmt.Println()
	fmt.Println("Slices of slice : ", slices, len(slices), cap(slices))
	fmt.Println("Slices of slice by make : ", makeSlices, len(makeSlices), cap(makeSlices))

	// 15. appending to a slice
	var oneDimensionSlice, newSlice []int

	fmt.Println()
	fmt.Println("Appending to a slice ")

	// first parameter of append is a slice of type T and rest are values of type T
	oneDimensionSlice = append(oneDimensionSlice, 0) // appending to a nil slice
	printIntSlice(oneDimensionSlice)

	newSlice = append(oneDimensionSlice, 1, 2, 3, 4) // slice grows as needed and can be appended with more than one element at a times
	printIntSlice(newSlice)
}

func printStringSlice(s []string) {
	fmt.Printf("Length = %v, Capacity = %v, Slice = %v\n", len(s), cap(s), s)
}

func printIntSlice(s []int) {
	fmt.Printf("Length = %v, Capacity = %v, Slice = %v\n", len(s), cap(s), s)
}
