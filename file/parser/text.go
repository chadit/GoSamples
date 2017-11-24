package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const textPath = "./text_files"

var lock = sync.RWMutex{}
var lineHashes = make(map[uint64]int)

func (h handler) processTextFiles() error {
	files, err := getFiles(textPath)
	if err != nil {
		return err
	}

	extractChannel := make(chan string)

	for _, file := range files {
		fp := textPath + "/" + file.Name()
		f, _ := os.Open(fp)

		buf := make([]byte, 32*1024) // define your buffer size here.

		for {
			n, err := f.Read(buf)

			if n > 0 {
				//lineHashes := make(map[uint64]int)
				for _, k := range strings.Split(string(buf[:n]), "\r\n") {
					extractChannel <- k
					// ll := len(k)
					// if ll == 0 {
					// 	continue
					// }

					// keywordParser(k)

					// if token := len(strings.Fields(k)); token > 0 {
					// 	tokensPerLine = append(tokensPerLine, token)
					// }

					// lineLengths = append(lineLengths, ll)
					// hl := hashFNV1a64(k)
					// if _, ok := lineHashes[hl]; ok {
					// 	dupCount++
					// } else {
					// 	lineHashes[hl] = 0
					// }
				}
			}

			if err == io.EOF {
				fmt.Println("io.EOF")
				break
			}
			if err != nil {
				log.Printf("read %d bytes: %v", n, err)
				break
			}
		}

	}

	return nil
}

func (h handler) parseLine() {
	var numMessages int64
	for l := range h.extractChannel {
		atomic.AddInt64(&numMessages, 1)
		go func(l string) {
			ll := len(l)
			if ll == 0 {
				return
			}

			keywordParser(l)

			if token := len(strings.Fields(l)); token > 0 {
				h.tokensPerLineChannel <- token
				//tokensPerLine = append(tokensPerLine, token)
			}
			h.lineLengthsChannel <- ll
			//	lineLengths = append(lineLengths, ll)
			hl := hashFNV1a64(l)
			if _, ok := lineHashes[hl]; ok {
				atomic.AddInt64(&dupCount, 1)
			} else {
				lock.Lock()
				lineHashes[hl] = 0
				lock.Unlock()
			}
			atomic.AddInt64(&numMessages, -1)
		}(l)
	}

	// checks to see if messages are still being processed before closing the channel
	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}

	close(h.lineLengthsChannel)
	close(h.tokensPerLineChannel)
}

func keywordParser(s string) {
	if s == "" {
		return
	}

	s = strings.ToLower(s)

	for _, keyword := range keywords {
		if strings.Contains(s, keyword) {
			kwmap[keyword]++
		}
	}
}
