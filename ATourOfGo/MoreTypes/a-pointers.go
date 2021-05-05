package main

import "fmt"

func main() {
	x, y := 4, 9

	// the & operator generates a pointer to its operand
	p := &x //referencing
	var q *int
	q = &y

	// the * operator denotes pointer's underlying value
	fmt.Println("x is read through pointer p = ", *p) // dereferencing or indirecting
	fmt.Println("y is read thorugh pointer q = ", *q)
	*p = *p * *p //	setting x through pointer p
	*q = y / 3   //	setting y thorugh pointer q

	fmt.Println("x is read through pointer p = ", *p) // dereferencing
	fmt.Println("y is read thorugh pointer q = ", *q)
}
