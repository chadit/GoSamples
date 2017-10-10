package main

import (
	"time"
)

// what you will see,  Hello and world will print in the order that the go routines were created,
// you should see a bunch of hellos then a bunch of worlds

func main() {
	// anonamous function
	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
		}
	}()

	// anonamous function
	go func() {
		for i := 0; i < 100; i++ {
			println("World")
		}
	}()

	// without a sleep the function will close out before the go routines has a chance to finish
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
