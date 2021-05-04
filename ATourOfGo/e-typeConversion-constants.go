package main

import (
	"fmt"
	"math"
)

// 15. Constants - declared with const keyword and not with var
// can be a character, string, boolean, or numeric values
// cannot be declared with := syntax
const globalConstantPi = 3.145

// 16. Numeric Constants - are high precision values
const (
	Big   = 1 << 100  // 1 is left shifted 100 times, i.e in binary 1 is followed by 100 zeros
	Small = Big >> 99 // shif Big right again by 99, so we end up with 1<<1 = 10 in binary i.e. 2
)

func main() {
	// 13. Type Conversion
	// The expression T(v) converts the value v to type T
	var i int = 45
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion
	// Removing float64 in math.Sqrt function call will cause error, as i won't be explicitly converted to float as f would expect
	var f float64 = math.Sqrt(float64(i))
	u := uint(f)

	// 14. Type Inference - When the right hand side of the declaration is typed, the new variable is of the same type
	var m float64
	j := m

	// But when the right hand side contains an untyped numeric constant, the new variable may be
	// an int, float64, or complex128 depending on the precision of the constant
	a := 42
	b := 3.142
	c := 3.5 + 0.5i

	fmt.Printf("Values and their types:\n")
	fmt.Printf("Type of %v is %T\n", i, i)
	fmt.Printf("Type of m is %T & as of j is %T\n", m, j)
	fmt.Printf("Type of %v is %T\n", f, f)
	fmt.Printf("Type of %v is %T\n", u, u)
	fmt.Printf("Type of a is %T, b is %T and c is %T\n", a, b, c)

	// Big cannot be prited due to overflow
	fmt.Printf("Type of Small is %T\n", Small)
}
