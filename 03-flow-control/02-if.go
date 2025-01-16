package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	// Note how the statement declaration can be made inside the if statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// `v` is dropped after the if/else blocks.
	return lim
}

func main() {
	defer fmt.Println("closing...") // Runs when the function it is held in returns
	// Can be stacked, and is executed in LIFO order

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// switch statements
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // If made conditionless, it just becomes a cleaner if/else chain
	case "darwin":
		fmt.Println("OS X.") // No break needed!
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
