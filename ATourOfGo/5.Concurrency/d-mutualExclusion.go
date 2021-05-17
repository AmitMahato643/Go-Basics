package main

import (
	"fmt"
	"sync"
	"time"
)

// 9. sync.Mutex : Mutual Exclusion
// Unlike communication between different goroutines through channel
// mutual exclusion is the method to ensure that the access to a shared variable
// is allowed to only a single goroutines at a time in order to avoid conflicts

// SafeCounter is safe to use concurrently.
// Go's sync.Mutex provides mutual exclusion with sync.Mutex having two methods Lock and Unlock
// these Lock and Unlock mehod helps us ensure that the shared variable is accessed only by one goroutine at a time
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock() // since only after the process has received the value we must unlock it, we use defer to achieve this
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		if i == 500 {
			// though 500 Inc goroutines has been instantiated at this point to run concurrently,
			// we can see from the output of this line that the variable actually incremented is always less than 500,
			// which indicates the actual completion of those numer of Inc goroutines
			fmt.Println(c.Value("somekey"))
		}
		go c.Inc("somekey")
	}

	fmt.Println("Wait for ", time.Second)
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
