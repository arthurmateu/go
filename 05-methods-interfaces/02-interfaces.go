package main

import "fmt"

type I interface { // Interfaces are sets of method signatures
	M() // Holds any value that implements these methods
	// Note that calling a method if the value held is nil will result on a runtime error
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

