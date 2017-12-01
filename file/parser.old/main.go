package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type handler struct {
	extractCh       chan string
	keywordCh       chan string
	lineLengthsCh   chan int
	tokensPerLineCh chan int
	doneCh          chan bool
	errorCh         chan error
	line            line
	key             key
}

type line struct {
	lock       sync.RWMutex
	lineHashes map[uint64]bool
	dupCount   int64
}

type key struct {
	lock     sync.RWMutex
	keywords []string
	kwmap    map[string]int
}

func main() {
	h := handler{
		extractCh:       make(chan string, 1000),
		keywordCh:       make(chan string, 1000),
		lineLengthsCh:   make(chan int, 1000),
		tokensPerLineCh: make(chan int, 1000),
		errorCh:         make(chan error, 100),
		doneCh:          make(chan bool),
		line:            line{lock: sync.RWMutex{}, lineHashes: make(map[uint64]bool)},
		key:             key{lock: sync.RWMutex{}, keywords: []string{}, kwmap: make(map[string]int)},
	}

	start := time.Now()
	fmt.Println("Parsing text")

	go h.processTextFiles()
	go h.parseLine()
	go h.keywordParser()

	llm := make(chan int, 5)
	lls := make(chan float64, 5)
	ttm := make(chan int, 5)
	tts := make(chan float64, 5)
	go h.lineStatParser(llm, lls)
	go h.tokenStatParser(ttm, tts)

	reportCh := make(chan string, 10)
	reportDoneCh := make(chan bool)
	errorDoneCh := make(chan bool)
	go logResults(reportCh, reportDoneCh, h.errorCh)
	go errorResults(h.errorCh, errorDoneCh)
	<-h.doneCh
	r := fmt.Sprintf("num dupes\t%d\n", h.line.dupCount) +
		fmt.Sprintf("med length\t%d\n", <-llm) +
		fmt.Sprintf("std length\t%.2f\n", <-lls) +
		fmt.Sprintf("med tokens\t%d\n", <-ttm) +
		fmt.Sprintf("std length\t%.2f\n", <-tts)

	for _, k := range h.key.keywords {
		r += fmt.Sprintf("keyword_%s\t%d\n", k, h.key.kwmap[k])
	}
	reportCh <- r
	close(reportCh)
	close(h.errorCh)
	<-reportDoneCh
	<-errorDoneCh
	fmt.Println(time.Since(start))
}

func logResults(reportCh chan string, reportDoneCh chan bool, errorCh chan error) {
	go func() {
		for {
			msg, ok := <-reportCh
			if ok {
				f, err := os.OpenFile("./results.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					errorCh <- err
					continue
				}
				defer f.Close()
				f.WriteString(msg)
			} else {
				reportDoneCh <- true
				break
			}
		}
	}()
}

func errorResults(errorCh chan error, errorDoneCh chan bool) {
	go func() {
		for {
			msg, ok := <-errorCh
			if ok {
				f, err := os.OpenFile("./errors.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if err != err {
					log.Println(err)
				}
				defer f.Close()
				f.WriteString(msg.Error())
			} else {
				errorDoneCh <- true
				break
			}
		}
	}()
}
