package main

import (
	"fmt"
	"os"
)

const (
	watchPath = "./test.txt"
	//hudredMB  = 100000000
	//oneMB = 1000000
	fiftyKB = 50000
)

func main() {
	fmt.Println("Hello")
	f, _ := os.OpenFile(watchPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
	f.Write(make([]byte, fiftyKB))
	f.Close()
}
