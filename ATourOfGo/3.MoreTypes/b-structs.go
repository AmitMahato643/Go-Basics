package main

import "fmt"

// A struct is a collection of fields
type Coordinate struct {
	x, y int
}

func main() {
	p1 := Coordinate{4, 4}
	fmt.Println("Coordinate p1 is ", p1)

	// struct fields can be accessed with dot (.) operator
	var p2 Coordinate
	p2.x = 5
	p2.y = 5

	fmt.Printf("x and y coordinate of p2 are %v and %v respectively\n", p2.x, p2.y)

	var (
		v1 = Coordinate{1, 2}  // has type Vertex
		v2 = Coordinate{x: 1}  // y:0 is implicit
		v3 = Coordinate{}      // x:0 and y:0
		p3 = &v3               // pointer to v3
		p4 = &Coordinate{1, 2} // has type *Vertex as & returns pointer to a struct value
	)

	// *p3 will dereference while p4 will simply display the referenced Coordinate
	fmt.Println(v1, *p3, p4, v2, v3)

}
