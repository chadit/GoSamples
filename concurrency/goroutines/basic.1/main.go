package main

import (
	"time"
)

// what you will see,  Hello and world will print alternating, first one will fire,
// then when it goes to sleep the second one will fire

func main() {

	godur, _ := time.ParseDuration("10ms")

	// anonamous function
	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(godur)
		}
	}()

	// anonamous function
	go func() {
		for i := 0; i < 100; i++ {
			println("World")
			time.Sleep(godur)
		}
	}()

	// without a sleep the function will close out before the go routines has a chance to finish
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
