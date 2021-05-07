package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // executes if c channel is ready
			{
				fmt.Println(x)
				x, y = y, x+y
			}
		case <-quit: // executes if quit channel is ready
			fmt.Println("quit")
			return
		default: // if non of the channel is ready select default case
			fmt.Println("waiting")
		}
	}
}

func main() {
	c := make(chan int, 3)
	quit := make(chan int)
	// comment out both the channel to see default channel in action
	go func() {
		for i := 0; i < 10; i++ {
			val, ok := <-c
			fmt.Println(val, ok)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
