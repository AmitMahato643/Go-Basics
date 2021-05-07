package main

import (
	"fmt"
	"io"
	"strings"
)

// 21. Readers
// the io.Reader interface has Read(b []byte) (n int, e err) method that populates b with x bytes of data where x is size of b,
// and returns n as number of byte  of data populated from the given stream and io.EOF error at the end of line

func main() {
	// here we are using Reader of strings
	str := strings.NewReader("Hello World")
	b := make([]byte, 8)
	for {
		// the Read method will fill b with max possible bytes of data and return the number of bytes populated and io.EOF error when completed
		n, err := str.Read(b)
		fmt.Printf("Read %v characters and b = %v\n", n, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // %q converts byte to specific character
		if err == io.EOF {
			break
		}
	}
}
