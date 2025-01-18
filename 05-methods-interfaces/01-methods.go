package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Since there's no classes, this is how you declare methods on types
// It could've been written as a regular function with 0 changes in functionality (Except on calling them)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This is a method with pointer receivers.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // For convenience, Go automatically interprets this as (&v).Scale(10)
	// Note that the opposite also happens (interpreting, for example, below as (*p).Abs()
	fmt.Println(v.Abs())
}
