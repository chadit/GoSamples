package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const dataFile = "./sample.txt"

func main() {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("marshal lines")
	str := string(data)
	lines := strings.Split(str, "\n")

	for i := range lines {
		fmt.Println(lines[i])
		info := strings.Split(lines[i], "|")
		fmt.Println(len(info))
	}

	//	fmt.Printf("\n%v\n", lines)

}
