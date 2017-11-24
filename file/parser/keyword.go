package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

const keywordPath = "./keywords"

func (h *handler) loadKeywords() {
	files, err := getFiles(keywordPath)
	if err != nil {
		h.errorCh <- err
		return
	}
	for _, file := range files {
		fp := keywordPath + "/" + file.Name()
		f, err := os.Open(fp)
		if err != nil {
			h.errorCh <- err
			continue
		}

		buf := make([]byte, 32*2024)

		for {
			n, err := f.Read(buf)

			if n > 0 {
				for _, k := range strings.Split(string(buf[:n]), "\n") {
					if k == "" {
						continue
					}
					k = strings.ToLower(k)
					h.key.lock.Lock()
					if _, ok := h.key.kwmap[k]; !ok {
						h.key.keywords = append(h.key.keywords, k)
						h.key.kwmap[k] = 0
					}
					h.key.lock.Unlock()

				}
			}

			if err == io.EOF {
				f.Close()
				break
			}
			if err != nil {
				log.Printf("read %d bytes: %v", n, err)
				f.Close()
				break
			}
		}

	}
	h.key.lock.Lock()
	sort.Strings(h.key.keywords)
	h.key.lock.Unlock()
}
