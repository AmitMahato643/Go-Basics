package main

import "fmt"

func main() {
	// 12. A defer statement defers the execution of a function until the surrounding function returns
	// The deffered call's argument are evaluated immediately but the function is not executed untill the surrounding function returns

	// Deferred function calls are pushed to stack, so it follows last-in-first-out order
	// Hence, following print statement goes into the bottom of the stack
	defer fmt.Printf("This will be printed just after this function returns\n")

	fmt.Println("Start")

	// Further print statements are added to the stack starting from 1 to 9 and hence 9 is at top
	for i := 0; i < 10; i++ {
		defer fmt.Printf("Number is %v\n", i)
	}

	fmt.Printf("This will be printed after Start\n")
	return
}
