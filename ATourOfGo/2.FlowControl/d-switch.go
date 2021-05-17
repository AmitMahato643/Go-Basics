package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Go Runs On ")

	// 9. Switch statement
	// like for loop and if statement, switch also has init statement and then the conditional
	// notice that there is no need of break statement at the end of each case block unlike as required in C, C++, Java, JavaScript, and PHP
	// Go does it implicitly
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "macOS":
		fmt.Println("macOS")
	case "windows":
		fmt.Println("Windows")
	}

	// 10. switch evaluation order
	// the evaluation order is from top to bottom so if case today matches then no other cases would be checked for
	fmt.Printf("\nWhen is Saturday? ")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Printf("Today\n")
	case today + 1:
		fmt.Printf("Tomorrow\n")
	case today + 2:
		fmt.Printf("In two days\n")
	default:
		fmt.Printf("Too far away\n")
	}

	// 11. switch without condition evaluates to switch true
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Print("Good Morning\n")
	case now.Hour() < 15:
		fmt.Print("Good Afternoon\n")
	case now.Hour() < 19:
		fmt.Print("Good Evening\n")
	case now.Hour() > 19:
		fmt.Print("Good Night\n")
	}

}
