package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// print out
// -- number of duplicate lines
// -----------------------------------------
// -- median length of lines
// -- -- slice of int, append length of each line, sort, if odd, get middle number
// -- -- if even, get center two items, add together then divide by 2 to get median
// -- standard deviation for length of lines -- https://www.easycalculation.com/statistics/standard-deviation.php
// -----------------------------------------
// -- median number of words (token) per line
// -- standard deviation for words (token) per line

/*
extract --
read each line, hash, add hash to map, if has exist inc dubLines
pass line to process channel
*/

/*
process --
pull line off channel,
append lineLength,
append # of tokens per line
kick off kerword search
*/

/*
signal when completed, so report can be generated
*/

type handler struct {
	extractChannel       chan string
	transformChannel     chan string
	lineLengthsChannel   chan int
	tokensPerLineChannel chan int
	doneChannel          chan bool
}

var (
	keywords = []string{}
	kwmap    = make(map[string]int)
	dupCount int64

//	lineLengths   = []int64{}
//	tokensPerLine = []int64{}
)

func main() {
	h := handler{
		extractChannel:       make(chan string),
		transformChannel:     make(chan string),
		lineLengthsChannel:   make(chan int),
		tokensPerLineChannel: make(chan int),
		doneChannel:          make(chan bool),
	}

	strart := time.Now()
	fmt.Println("Hello")
	if err := loadKeywords(); err != nil {
		fmt.Printf("failed to load keywords : %v\n", err)
		return
	}

	if err := h.processTextFiles(); err != nil {
		fmt.Printf("failed to load text : %v\n", err)
		return
	}
	lm, ls := calcstats(h.lineLengthsChannel)
	tm, ts := calcstats(h.tokensPerLineChannel)
	fmt.Println("Results:\n---------------------------------------")
	fmt.Printf("num dupes  \t\t| %d\n", dupCount)
	fmt.Printf("line med length \t| %d\n", lm)
	fmt.Printf("line std length \t| %.2f\n", ls)
	fmt.Printf("token med length \t| %d\n", tm)
	fmt.Printf("token std length \t| %.2f\n", ts)

	for _, k := range keywords {
		fmt.Printf("%s%s| %d\n", k, setTabs(k), kwmap[k])
	}

	fmt.Println(time.Since(strart))
}

func setTabs(s string) string {
	c := int(math.Floor(float64(len(s) / 5)))
	s1 := "\t\t\t"
	for i := 0; i < c; i++ {
		s1 = strings.Replace(s1, "\t", "", 1)
	}
	return s1
}
