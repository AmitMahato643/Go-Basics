package main

import "fmt"

type Stringer interface {
	String() string
}

type IPAddr [4]byte

// 17. Stringer - One of the most ubiquitous (present, appearing, found) interfaces is Stringer defined by the fmt package
// A Stringer is a type that describes itself as a string
// fmt package and many others look for this interface to print values
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3]) // returns a string
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
