package main

import (
	"fmt"
	"runtime"
	"sync"
)

// basic mutex, can be slow, due to async to sync calls
func main() {
	mutex := new(sync.Mutex)
	// can run accross up to 4 cores
	runtime.GOMAXPROCS(4)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex.Lock()
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				mutex.Unlock()
			}()
		}
	}

	// waits for a key to be pressed
	fmt.Scanln()
}
