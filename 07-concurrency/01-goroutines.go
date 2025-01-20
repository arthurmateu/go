package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // Select blocks until this runs.
		// Chooses one at random if multiple are ready.
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func explosion() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: // If no other case is ready, this one will be run
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	go say("world") // goroutines are lightweight threads
	say("hello")

	time.Sleep(5000)
	println("===")

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

	time.Sleep(5000)
	println("===")

	explosion()
}
