package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/* Fibonacci Closure
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		prev := a
		a, b = b, a+b
		return prev
	}
}
*/
