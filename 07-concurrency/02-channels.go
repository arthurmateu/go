package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // channels must be created before
	// Note that they can be buffered, by passing a second argument (being its length) on make()
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c, and assign value to variables
	// again, receivers can test whether or not channels are open/closed like this:
	// v, ok := <-ch

	fmt.Println(x, y, x+y)

	close(c) // Signalizes that no more values will be sent.
	// Not usually necessary, like in the case of files.
	// NOTE: only the sender should close the channel.
}
