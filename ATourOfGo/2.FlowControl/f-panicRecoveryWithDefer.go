package main

import "fmt"

func DoMath(a int, b int) {
	var sum, diff, prod, quot int

	// If b=0 this function will return a panic if not handled and the program will not terminate normally
	// handle the panic that will be caused when b=0; divide by zero panic occurs
	// choose alternate flow after panic is occured than the normal one
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error occurred while performing math operations on given data, error : %v\n", r)
			fmt.Printf("Skipping the division\n")
			fmt.Printf("Sum = %v, Difference = %v, Product = %v\n", a+b, a-b, a*b)
		}
	}()

	sum, diff, prod, quot = a+b, a-b, a*b, a/b

	fmt.Printf("Sum = %v, Difference = %v, Product = %v, Quotient = %v\n", sum, diff, prod, quot)
}

func main() {
	var a, b int = 4, 2 // will not cause panic and program flow will be normal
	var m, n int = 4, 0 // will cause panic and program flow will not be normal
	DoMath(a, b)
	DoMath(m, n)
}
