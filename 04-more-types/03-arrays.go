package main

import (
	"fmt"
	// "golang.org/x/tour/pic"
)

func main() {
	var a [2]string // Note that it can't be resized
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // Initialized on one-line
	fmt.Println(primes)
	var s []int = primes[1:4] // Slices. Note that it references to the original array, so any changes made to it affect the referenced array.

	s = append(s, 17, 19) // Appending can take 1+ elements
	fmt.Println(s)
	fmt.Println(primes) // ...see how it also affects the original one, writing over the other 2 values?

	for index, value := range s { // Equivalent to enumerate()
		fmt.Printf("Value at index %d: %d\n", index, value)
	}

	// pic.Show(Pic)
}

/*
func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := range p {
		p[y] = make([]uint8, dx)
		for x := range p[y] {
			p[y][x] = uint8(x*y)
		}
	}
	return p
}
*/
