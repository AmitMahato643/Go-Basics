// 1. Goroutines : A goroutine is a lightweight thread managed by Go Runtime
// go f(x,y) starts a new goroutine, evaluation of f, x, y happen in current routine while execution of f in new routine
// since goroutines run in same address space, access to shared memory must be synchronized

// 2. Channel : helps to achieve communication between goroutines
// send and receive values with the channel operator <-
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // only the sender can close the channel; sending on a closed channel will cause a panic
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// spread the numbers to concurrently process the sum of slice elements
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	// finally add the results of concurrent goroutines for final value of sum
	fmt.Println(x, y, x+y)

	// 3. Buffered channel; predefine the size of elements that can be buffered into the channel
	// Before anymore elements than the size of buffer to are added, previous value must be taken out first
	// inserting more than specified buffer size will cause deadlock - overflow
	bc := make(chan int, 2)
	bc <- 1
	bc <- 2
	fmt.Printf("First two values in channel : %v & %v\n", <-bc, <-bc)
	// channel is completely empy now
	bc <- 3
	bc <- 4
	// channel is again full
	fmt.Printf("Third values in channel : %v\n", <-bc)
	// channel has one value already
	bc <- 3 // channel is again full
	// bc <- 4 // this will cause overflow

	// 4. Range and Close
	bufferSize := 7
	c = make(chan int, bufferSize)
	go fibonacci(cap(c), c) // will generate first bufferSize number of of fibonacci numbers

	// Using Range
	for i := range c { // it doesn't return index,value tuple instead channel only return value
		fmt.Println(i)
	}

	// Demonstration of using close, ok is false after channel is closed
	fmt.Println()
	for i := 0; i < cap(c); i++ {
		v, ok := <-c
		if ok {
			fmt.Printf("Value received in %v attempt : %v\n", i+1, v)
		} else {
			fmt.Printf("%v. Channel already closed\n", i+1)
		}
	}
}
