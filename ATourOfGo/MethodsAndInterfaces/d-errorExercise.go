package main

import (
	"fmt"
	"math"
)

// 19. 20. Errors
// Took a while to get some level of understanding
// ref: https://tour.golang.org/methods/20

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// simply using e in fmt.Sprintf(...,e) cause it to go in an infinite loop so convert it to float64
	// possible reason is that, the fmt.Sprintf receives value of type ErrNegativeSqrt and again founds Error method in it
	// and hence this method is again called and so on
	return fmt.Sprintf("Error, negative value %v received", float64(e))
}

func (e ErrNegativeSqrt) String() string {
	return fmt.Sprintf("%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(ErrNegativeSqrt(4))
}
