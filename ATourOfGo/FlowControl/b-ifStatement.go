package main

import "fmt"

func main() {
	name := "Amit"
	// 5. the if statement, lik for loop doesn't require paranthese wrapping the conditional expression
	// 6. a short statement can be used just like the init statement in for loop, which is executed befor the conditional expression
	// variable declared in the statement are within the scope of if-else block
	if length, longName := len(name), "Amit Mahato"; length < 7 {
		fmt.Printf("The name %v is a short name\n", name)
	} else if longNameLen := len(longName); longNameLen > 7 {
		fmt.Printf("The name %v is a long name\n", longName)
	}

	// longName and longNameLen cannot be used here, uncomment below print statement to see error
	// fmt.Printf("Long Name : %v",longName)
}
