package main

import (
	"fmt"
	"strings"
)

// 1. Metods: Go doesn't have classes but a type can have methods associated with them
type Student struct {
	name, faculty     string
	rollNumber, batch int
}

// A method is simply a function with special receiver argument
// The receiver appears in its own argument list between the func keyword and the method name
// Here, the getFullRollNum method has a receiver of type Student named s
// the named type receiver can be considered as replacement for 'this' and 'self' keywords as used in other programming languages
func (s Student) getFullRollNum() string {
	b := fmt.Sprint(s.batch)[1:] // convert batch to string and then drop element at index 0,i.e. drop its first character
	return strings.Join([]string{b, s.faculty, fmt.Sprint(s.rollNumber)}, "")
}

// 2. Methods are functions: a method is just a function with a special receiver argument
// the below written function will behave same as above with no change in its functionality
func getFullRollNumber(s Student) string {
	b := fmt.Sprint(s.batch)[1:] // convert batch to string and then drop element at index 0,i.e. drop its first character
	return strings.Join([]string{b, s.faculty, fmt.Sprint(s.rollNumber)}, "")
}

// 3. Methods Continued - Methods can be declared on non-sturct types too
type MyFloat float64

// Here MyFloat which isn't a struct has Abs method implemented on it
func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return MyFloat(-f)
	} else {
		return MyFloat(f)
	}
}

// Methods can only be declared with a receiver whose type is defined in the same package as the method
// Mthods cannot be declared with a receiver whose type is defined in another package (including built-in types int)
// try this, you cannot access Abs method on int type variables
// error here, invalid receiver int (basic or unnamed type); here its basic type : int
// func (i int) Abs() int {
// 	if i < 0 {
// 		return int(-i)
// 	} else {
// 		return int(i)
// 	}
// }

type Rectangle struct {
	length, breadth MyFloat
}

// 8.		Choosing a value or pointer receiver - Two Reasons
// 8.1	To modify the value that the reciever points to
// 8.2	To avoid copying the value on each method call. This can be more efficient if the receiver is a large struct

// 8.	Ex. The Area method can too take pointer receiver although it need not modify
// 				the values that the receiver points to, for increasing efficiency
func (r Rectangle) Area() float64 {
	return float64(r.length.Abs()) * float64(r.breadth.Abs())
}

// 4. Pointer Receiver : Methods can be declared with pointer receivers
// Receiver type is of type *T and T cannot itself be a pointer such as *int
// Methods with pointer receiver can modify the value to which the receiver points (as shown in scale here)

// With a value receiver, the operation is done on a copy of original Rectangle, just like done with any other arguments of a function
// Try removing * from receiver argument to see the effect

func (r *Rectangle) scale(s float64) {
	r.length = MyFloat(s * float64(r.length))
	r.breadth = MyFloat(s * float64(r.breadth))
}

// 5. Pointer and functions : If the same has to be done with normal function
// it would take pointer as argument and hence while calling reference has to be sent as params
func scaleRectangle(r *Rectangle, s float64) {
	r.length = MyFloat(s * float64(r.length))
	r.breadth = MyFloat(s * float64(r.breadth))
}

func main() {
	student := Student{
		name:       "Amit Mahato",
		batch:      2073,
		faculty:    "BCT",
		rollNumber: 505,
	}

	// call a method
	fmt.Printf("Hello %v, I found your complete roll number to be %v\n", student.name, student.getFullRollNum())

	// call a function
	fmt.Printf("Hello %v, I found your complete roll number to be %v\n", student.name, getFullRollNumber(student))

	rectangle := Rectangle{
		length:  MyFloat(-10),
		breadth: MyFloat(-5),
	}
	fmt.Printf("Area of rectangle having lenght %v and breadth %v = %v\n", rectangle.length.Abs(), rectangle.breadth.Abs(), rectangle.Area())

	scaleBy := 2.5
	// 6. Methods and Pointer indirection - methods with pointer receiver take either a value or a pointer as the receiver
	// As convinience, Go interprets the statement rectangle.scale(scaleBy) as (&rectangle).scale(scaleBy) since the scale method has a pointer receiver
	rectangle.scale(scaleBy)
	fmt.Printf("Rectangle after scaling by %v times : %v with area %v\n", scaleBy, rectangle, rectangle.Area())

	scaleBy = 0.5
	// Function with pointer as argument must take a pointer
	// scaleRectangle(rectangle, scaleBy) // Compile error!
	// Equivalent thing happen in reverse direction, function with value as argument must take value and not reference, while both are OK for methods
	scaleRectangle(&rectangle, scaleBy) // OK
	fmt.Printf("Again, Rectangle after scaling by %v times : %v with area %v\n", scaleBy, rectangle, rectangle.Area())

	return
}
