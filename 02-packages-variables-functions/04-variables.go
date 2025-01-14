package main

import "fmt"

func main() {
	const greet string = "hello"
	var i, j int = 1, 2 // Multiple assignments
	k := 3 // Inferred type. Doesn't work with constants
	c, python, java := true, false, "no!" // Multiple assignments with inferred types

	fmt.Println(i, j, k, c, python, java)
}

