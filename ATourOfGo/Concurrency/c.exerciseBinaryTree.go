// 8. Exercise

// Steps
// 1. Implement the Walk function.

// 2. Test the Walk function.
// The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
// Create a new channel ch and kick off the walker:
// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

// 4. Test the Same function.
// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

package main

import (
	"fmt"
	"math"
	"strconv"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) /*[]int*/ { // same could have been done without channel
	// var s []int
	var temp, count int
	var start bool
	temp = 0
	for _, str := range t.String() {
		n, err := strconv.Atoi(string(str))
		if err == nil {
			temp = temp*int(math.Pow10(count)) + n
			count += 1
			start = true
		} else {
			if start {
				// s = append(s, temp)
				ch <- temp // send the value of temp through the channel
				start = false
				temp = 0
				count = 0
			}
		}
	}
	close(ch) // only the sender should close the channel
	// return s
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	same := true
	c1 := make(chan int) // channel creation for data sharing between different goroutines
	c2 := make(chan int)

	// concurrently run two instances of Walk function
	go Walk(t1, c1) // go Walk is evaluated by current routine but Walk is executed in new goroutine
	go Walk(t2, c2) // go Walk is evaluated by current routine but Walk is executed in new goroutine

	for {
		t1val, ok1 := <-c1 // grab value received from channel and boolean for channel is closed or not info
		t2val, ok2 := <-c2

		if ok1 {
			fmt.Printf("c1 = %v ", t1val)
		}
		if ok2 {
			fmt.Printf("c2 = %v ", t2val)
		}
		if ok1 == false && ok2 == false {
			break // break if both the channels are closed
		}
		fmt.Println()
		if t1val != t2val {
			same = false // if values mismatch they are not same
		}

	}
	return same
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println("Is same : ", Same(t1, t2))
}
