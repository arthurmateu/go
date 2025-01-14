package main

import "fmt"

func swap(x, y string) (string, string) { // Type comes after
	return y, x
}

/*
 *	Naked function example:
 *
 *	func split(sum int) (x, y int) { <- Note how the result variables are named
 *		x = sum * 4 / 9
 *		y = sum - x
 *		return <- That way, you don't need to include them here
 *	}
 *
 *  Avoid using them, though.
*/

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

