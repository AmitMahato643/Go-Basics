package main

import "fmt"

// a function can accept another function as an argument
func display(adder func(int, int) int, subtractor func(int, int) int) {
	minDiceValue := 1
	maxDiceValue := 6
	maxPlusMinDice := adder(maxDiceValue, minDiceValue)
	maxMinusMinDice := subtractor(maxPlusMinDice, minDiceValue)

	fmt.Printf("The sum of minimum and maximum values of a Dice is: %v\nThe difference between maximum and minimum values of a Dice is: %v\n", maxPlusMinDice, maxMinusMinDice)
}

func main() {
	// a function can be assigned to a variable
	adder := func(a, b int) int {
		return a + b
	}

	subtractor := func(a, b int) int {
		return a - b
	}

	a := 2
	b := 5
	fmt.Printf("Sum of %v and %v = %v\n", b, a, adder(a, b))
	fmt.Printf("Difference between %v and %v = %v\n", b, a, subtractor(b, a))

	fmt.Println("********** Sum and difference between min and max value in a Dice *********")
	display(adder, subtractor) // display function accepts adder and subtractor function as its parameters

	// 25. Function closure
	// A Go function may be a closure. A closure is a function value that references
	//  variables from outside its body. Th. function may access and assign to the
	// referenced variables; int this sense the function is "bound" to the variables.
	// Best example of closure is shown in the exercise for function closure in the next file- The Fibonacci Closure
}
