package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// accomplished via channels
func main() {
	// can run accross up to 4 cores
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				fmt.Println(msg)
				f, _ := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				//f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				f.WriteString(time.Now().Format(time.RFC3339) + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				logCh <- fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				<-mutex
			}()
		}
	}

	// waits for a key to be pressed
	fmt.Scanln()
}
