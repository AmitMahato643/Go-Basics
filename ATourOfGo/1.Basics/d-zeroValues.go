// 12. Default values of any type / Zero values
package main

import "fmt"

func main() {
	var number int
	var decimal float64
	var name string
	var userActive bool

	fmt.Printf("Default value of different types:\n%T is %v\n%T is  %v\n%T is  %v\n%T is  %q\n", number, number, decimal, decimal, userActive, userActive, name, name)
}
