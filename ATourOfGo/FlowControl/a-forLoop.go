package main

import "fmt"

func main() {
	// 1. The for loop
	// only one looping construct in Go : the for loop
	// for loop has three components
	// 1. init statement : executed before the first iteration
	// 2. condition expression : checked before each iteration
	// 3. the post statement : executed at the end of each iteration
	// Note: Unlike C, Java or Javascript, these three components are not wrapped around with parantheses
	// and the curly braces {} are always required
	for i := 0; i < 5; i++ {
		fmt.Printf("Hi for %vth time\n", i+1)
	}

	// 2. for loop continued - the init and post statements are optional
	sum := 1
	for sum < 100 {
		fmt.Printf("Sum = %v\n", sum)
		sum += sum
	}
	fmt.Printf("Total sum is : %v\n", sum)

	// 3. Go's while loop is the for loop
	// after drooing semicolons, for loop can cover the while loop
	lessThanEqualToTen := 1
	for lessThanEqualToTen < 10 {
		lessThanEqualToTen += 1
	}
	fmt.Printf("lessThanEqualToTen = %v\n", lessThanEqualToTen)

	// 4. Forever - The infinite for loop
	for {
		fmt.Printf("I won't die on my own. Kill me please, with Ctrl + C\n")
	}
}
