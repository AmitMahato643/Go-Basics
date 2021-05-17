package main

import "fmt"

// MyFloat
type MyFloat float64

func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return MyFloat(-f)
	} else {
		return f
	}
}

// 9. Interface - defined as a set of method signatures
// A value of interface type can hold any value that implements those methods

// Shape interface which would implement
// 1. Self method taking nothing & returning map containing the properties of underlying type (struct in our case)
// 2. Scale method taking float64 argument & returning nothing
// 3.	Area method taking nothing & returning float64 value
// Notice both should have either pointer receiver or value receiver but not a mixture of them
type Shape interface {
	Self() map[string]MyFloat
	Scale(float64)
	Area() float64
}

// 14. The empty interface - the interface type that specifies zero type; holds any type (every type implements at least zero method)
// i of type empty interface
func describe(i interface{}) {
	// an interface is a tuple of concrete value and concrete type
	fmt.Printf("Interface Description : (%v, %T)\n", i, i)
}

// 10. Interfaces are implemented implicitly
// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implement" keyword
// Implicit interfaces decouples the definition of interface from its implementation, which could then appear in any package.

// Circle - Implementing Shape Interface
type Circle struct {
	radius MyFloat
}

// Implementation of these methods: Self, Scale and Area, means Circle implements interface Shape
// but its not needed to explicity declare that it does
func (c *Circle) Self() map[string]MyFloat {
	m := make(map[string]MyFloat)
	if c == nil {
		fmt.Println("<nil>")
		return m
	}

	m["radius"] = c.radius
	return m
}

func (c *Circle) Area() float64 {
	Pi := 3.1415

	if c == nil {
		fmt.Println("<nil>")
		return 0
	}

	return Pi * float64(c.radius.Abs())
}

func (c *Circle) Scale(s float64) {

	if c == nil {
		fmt.Println("<nil>")
		return
	}

	c.radius = MyFloat(float64(c.radius) * s)
}

// Rectangle - Implementing Shape Interface
type Rectangle struct {
	length, breadth MyFloat
}

func (r *Rectangle) Self() map[string]MyFloat {
	m := make(map[string]MyFloat)
	if r == nil {
		fmt.Println("<nil>")
		return m
	}
	m["length"] = r.length
	m["breadth"] = r.breadth
	return m
}

func (r *Rectangle) Area() float64 {
	if r == nil {
		fmt.Println("<nil>")
		return 0
	}
	return float64(r.length.Abs()) * float64(r.breadth.Abs())
}

func (r *Rectangle) Scale(s float64) {
	if r == nil {
		fmt.Println("<nil>")
		return
	}
	r.length = MyFloat(float64(r.length) * s)
	r.breadth = MyFloat(float64(r.breadth) * s)
}

func main() {
	var shape Shape
	scaleBy := 2.5

	c := Circle{
		radius: MyFloat(5),
	}
	// not using & here can cause error, since Scale and Area have pointer receiver and not value receiver
	// i.e. they are defined on *Circle and not on Circle
	shape = &c
	shape.Scale(scaleBy)

	fmt.Println()
	describe(shape)
	fmt.Printf("Area of shape, %T %v scaled by scaling factor of %v is %v\n", c, shape.Self(), scaleBy, shape.Area())

	r := Rectangle{
		length:  -10,
		breadth: -5,
	}

	shape = &r
	shape.Scale(scaleBy)

	fmt.Println()
	describe(shape)
	fmt.Printf("Area of shape, %T %v scaled by scaling factor of %v is %v\n", r, shape.Self(), scaleBy, shape.Area())

	// 12. Interface value with nil underlying value:
	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver
	var rNil *Rectangle // nil rNil Rectangle
	// Notice that the interface value that holds a nil concrete value is itself non-nil
	shape = rNil // non-nil shape
	fmt.Println()
	describe(shape)
	shape.Self()

	// 13. A nil interface
	// A nil interface value holds neither concrete type nor a value
	// Calling the method on nil interface causes runtime error
	// as there is no type inside the interface tuple to indicate concrete method to call
	var nilShape Shape
	fmt.Println()
	describe(nilShape)
	// nilShape.Scale(2) // will cause runtime error

	// 15. Type Assertions
	// Type assertion provides access to interface value's underlying concrete value
	// t = i.(T) ; t will be concrete value of type T; if not the statement will trigger a panic
	// To test whether a interface value holds specific type, a type assertion returns two values
	// the underlying value and a boolean value that reports whether the assertion succeed
	var recInterface interface{} = Rectangle{length: 5, breadth: 4}
	rVal, ok := recInterface.(Rectangle) // its like accessing properties of specific class
	if ok {
		describe(recInterface)
		fmt.Println("Value = ", rVal)
	}

	cVal, ok := recInterface.(Circle)
	fmt.Println("Value = ", cVal)

	// f := recInterface.(float64) // panic - in order to avoid panic, grab boolean value as well i.e. f, ok = recInterface(float64)
	// fmt.Println("Value = ", f)

	// 16. Type Switches demo
	fmt.Println()
	findType(recInterface)
	findType(Circle{5})
	findType(recInterface.(Rectangle).breadth)
}

func findType(i interface{}) {
	// 16. Type Switches
	switch v := i.(type) {
	case Rectangle:
		fmt.Println("Received rectangle is ", v)
	case Circle:
		fmt.Println("Received circle  is ", v)
	case MyFloat:
		fmt.Println("Received MyFloat is ", v)
	default:
		fmt.Println("None of user defined types received")
	}

}
