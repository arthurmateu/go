package main

import "fmt"

func main() {
	sum := 0
	// Regular for loop. You can ignore the init and post statements
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// This is a while
	for sum < 1000 {
		sum += sum
	}
	// This, an infinite loop
	for {
	}
}

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v iter -> %v\n", i+1, z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}
