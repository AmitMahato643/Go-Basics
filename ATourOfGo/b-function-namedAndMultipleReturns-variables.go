// 4. Functions
package main

import "fmt"

// 4. Functions
// A function can take zero or more arguments, here, add takes two arguments of type int
// Unlike C, type comes after variable name
// Notice the return type comes after arguments and before function definition starts
// here return type is int
func add(a int, b int) int {
	return a + b
}

// 5. Function Continued
// if two or more consecutive named function parameters share a type, type from all but the last can be omitted
func subtract(a, b int) int {
	return a - b
}

// 6. Multiple Results
// 7. Named Return values : A Go's return value may be named, if so, they are treated as variables defined at top of the funcion
func AddSubtract(a int, b int) (addition, subtraction int) {
	addition = add(a, b)
	subtraction = subtract(a, b)

	// 7. A return statement without argument returns the named return values
	// 7. This is known as naked return; naked return should be used only in short functions
	// it better to do return addition, subtraction
	return
}

var aGlobalVariable string = "This is a global variable"

func main() {

	// 8. Variables - A var statement declares a list of variables; as in function argument list, the type is present at last
	// A var statement can be at package (aGlobalVariable demonstrated above) or function level (x, y variables below)
	// 9. Variables with initializer - a var declaration can include initializers, one per variable
	var x, y int = 5, 4
	fmt.Printf("Sum of %d and %d is %d.\n", x, y, add(x, y))
	fmt.Printf("Subtraction of %d and %d is %d.\n", x, y, subtract(x, y))

	// 9. if the initializer is present, the type can be omitted, variable takes the type of initializer
	// 10. Short variable declarations
	// 10.1 Inside a function, the := short assignment can be used in place of var declaration with implicit type
	// 10.2 Outside a function, every statement begins with a keyword (var, func, etc) and hence := construct is not available
	addition, subtraction := AddSubtract(x, y)
	fmt.Printf("Demonstrates multiple result of a function, addition = %d, subtraction = %d\n", addition, subtraction)
}
