package main

import "fmt"

var (
	message = "Hello World"
)

func main() {
	fmt.Println(message)
}

func init() {
	message = "Hello GO!!!!"
}
