// 11. Basic Types in Go
package main

import "fmt"

func main() {
	var userActive bool = true
	var name string = "amit"
	var about string = `I am Amit Mahato.
	Completed my BE from Pulchowk Engineering Campus
	This is an example of multiline string`
	var rollNum int = 505
	var grade float64 = 3.4
	complex := 3 + 4i // complex128
	num := byte('a')

	fmt.Printf("Type is %T, value is %v\n", userActive, userActive)
	fmt.Printf("Type is %T, value is %v\n", name, name)
	fmt.Printf("Type is %T, value is %v\n", about, about)
	fmt.Printf("Type is %T, value is %v\n", rollNum, rollNum)
	fmt.Printf("Type is %T, value is %v\n", grade, grade)
	fmt.Printf("Type is %T, value is %v\n", complex, complex)
	fmt.Printf("Type is %T, value is %v\n", num, num)
}
