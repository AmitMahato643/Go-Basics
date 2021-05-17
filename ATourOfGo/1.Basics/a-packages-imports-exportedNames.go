// 1. Packages
// 1. Every Go program is made of packages & program starts in package main
package main

// 1. Packages
// packages fmt and math/rand are imported here
// by convention, a package name is same as the last element of the import path,
// math/rand package comprise file that begin with the statement package rand

// 2. Imports
// This works too, but linter for Go on saving the file automatically follows "factored" import statement
// multiple import statement
// import "fmt"
// import "math/rand"

// good style of import - "factored" import statement
import (
	"fmt"
	"math/rand"
)

// 1. The program starts from the main function
func main() {
	// 1. the environment in which these programs are executed are deterministic and hence rand.Int() will return same value everytime
	// 3. Exported names: In Go, a name is exported if it begins with capital letter, Println is exported but not println
	fmt.Println("Favorite number : ", rand.Intn(54))
}
