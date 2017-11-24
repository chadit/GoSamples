package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

const textPath = "./text_files"

func (h *handler) processTextFiles() {
	files, err := getFiles(textPath)
	if err != nil {
		h.errorCh <- err
		return
	}

	for _, file := range files {
		fp := textPath + "/" + file.Name()
		f, err := os.Open(fp)
		if err != nil {
			h.errorCh <- err
			return
		}

		buf := make([]byte, 32*1024)

		for {
			n, err := f.Read(buf)

			if n > 0 {
				for _, k := range strings.Split(string(buf[:n]), "\r\n") {
					if len(k) == 0 {
						continue
					}
					h.extractCh <- k
				}
			}

			if err == io.EOF {
				f.Close()
				break
			}
			if err != nil {
				h.errorCh <- fmt.Errorf("read %d bytes: %v", n, err)
				f.Close()
				break
			}
		}
	}
	close(h.extractCh)
}

func (h *handler) parseLine() {
	var numMessages int64
	for l := range h.extractCh {
		atomic.AddInt64(&numMessages, 1)
		go func(l string) {
			ll := len(l)
			if ll > 0 {
				h.keywordCh <- l
				if token := len(strings.Fields(l)); token > 0 {
					h.tokensPerLineCh <- token
				}
				h.lineLengthsCh <- ll
				hl := hashFNV1a64(l)
				h.line.lock.Lock()
				if _, ok := h.line.lineHashes[hl]; ok {
					if h.line.dupCount == 0 {
						atomic.AddInt64(&h.line.dupCount, 1)
					}
					atomic.AddInt64(&h.line.dupCount, 1)
				} else {
					h.line.lineHashes[hl] = true
				}
				h.line.lock.Unlock()
			}
			atomic.AddInt64(&numMessages, -1)
		}(l)
	}

	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}

	close(h.lineLengthsCh)
	close(h.tokensPerLineCh)
	close(h.keywordCh)
}

func (h *handler) lineStatParser(llm chan int, lls chan float64) {
	d := []int{}
	for n := range h.lineLengthsCh {
		d = append(d, n)
	}
	llm <- calcMedian(d)
	lls <- calcSD(d)
}

func (h *handler) tokenStatParser(ttm chan int, tts chan float64) {
	d := []int{}
	for n := range h.tokensPerLineCh {
		d = append(d, n)
	}
	ttm <- calcMedian(d)
	tts <- calcSD(d)
}

func (h *handler) keywordParser() {
	var numMessages int64
	h.loadKeywords()

	for l := range h.keywordCh {
		atomic.AddInt64(&numMessages, 1)
		go func(l string, h *handler) {
			l = strings.ToLower(l)

			h.key.lock.Lock()
			for _, keyword := range h.key.keywords {
				if strings.Contains(l, keyword) {
					h.key.kwmap[keyword]++
				}
			}
			h.key.lock.Unlock()
			atomic.AddInt64(&numMessages, -1)
		}(l, h)
	}

	for numMessages > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	h.doneCh <- true
}
