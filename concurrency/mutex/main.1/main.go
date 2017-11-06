package main

import (
	"fmt"
	"runtime"
)

// accomplished via channels
func main() {
	// can run accross up to 4 cores
	runtime.GOMAXPROCS(4)
	mutex := make(chan bool, 1)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-mutex
			}()
		}
	}

	// waits for a key to be pressed
	fmt.Scanln()
}
