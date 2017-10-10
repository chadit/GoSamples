package main

import (
	"runtime"
	"time"
)

// what you will see,  Hello and world will print in a random order depending on which thread is asleep
// and how long it takes to print the line out... this is an example of concurrancy and parallelism

func main() {
	godur, _ := time.ParseDuration("10ms")
	// we are telling our application it can only use 2 processes
	runtime.GOMAXPROCS(2)

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
