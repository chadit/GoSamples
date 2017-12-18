package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/profile"
)

type handler struct {
	Lines
	Keys
	Stats Stats
}
type Stats struct {
	LineDupCount int64
	LineMedian   int
	LineStDev    float64
	TokenMedian  int
	TokenStDev   float64
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	h := handler{
		Lines: Lines{},
		Keys:  Keys{},
	}

	start := time.Now()
	fmt.Println("Parsing text")
	if err := h.LoadKeywords(); err != nil {
		log.Println(err)
	}
	if err := h.ProcessTextFiles(); err != nil {
		log.Println(err)
	}
	h.StatsParser()

	r := fmt.Sprintf("num dupes\t%d\n", h.Stats.LineDupCount) +
		fmt.Sprintf("med length\t%d\n", h.Stats.LineMedian) +
		fmt.Sprintf("std length\t%.2f\n", h.Stats.LineStDev) +
		fmt.Sprintf("med tokens\t%d\n", h.Stats.TokenMedian) +
		fmt.Sprintf("std length\t%.2f\n", h.Stats.TokenStDev)

	for _, k := range h.Keys {
		r += fmt.Sprintf("keyword_%s\t%d\n", k.Name, k.Count)
	}
	logResults(r)
	fmt.Println(time.Since(start))
}

func logResults(r string) {
	f, err := os.OpenFile("./results.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString(r)
}
