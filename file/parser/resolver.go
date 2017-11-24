package main

import (
	"hash/fnv"
	"math"
	"os"
	"sort"
)

func stringExist(kw string, kws []string) bool {
	for i := range kws {
		if kws[i] == kw {
			return true
		}
	}
	return false
}

func hashFNV1a64(s string) uint64 {
	h := fnv.New64a()
	if _, err := h.Write([]byte(s)); err != nil {
		return 0
	}
	return h.Sum64()
}

func getFiles(s string) ([]os.FileInfo, error) {
	d, err := os.Open(s)
	if err != nil {
		return []os.FileInfo{}, err
	}

	// get all the files in the director
	return d.Readdir(-1)
}

func calcstats(c chan int) (int, float64) {
	d := []int{}
	for n := range c {
		d = append(d, n)
	}
	m := calcMedian(d)
	sd := calcSD(d)
	return m, sd
}

func calcMedian(d []int) int {
	n := len(d)
	if n == 0 {
		return 0
	}
	sort.Ints(d)
	if n%2 == 0 {
		return (d[n/2] + d[(n/2)-1]) / 2
	}
	return d[n/2]
}

func calcSD(d []int) float64 {
	l := len(d)
	if l == 0 {
		return 0
	}

	m := calMean(d, l)
	if m == 0 {
		return 0
	}

	sum := 0
	for _, d1 := range d {
		sum = sum + square((d1-m), 2)
	}
	return math.Sqrt(float64(sum / l))
}

func calMean(d []int, l int) int {
	n1 := 0
	for _, d1 := range d {
		n1 = n1 + d1
	}

	return n1 / l
}

func square(n1, pow int) int {
	return int(math.Pow(float64(n1), float64(pow)))
}
