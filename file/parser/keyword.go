package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

const keywordPath = "./keywords"

func loadKeywords() error {
	files, err := getFiles(keywordPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		fp := keywordPath + "/" + file.Name()
		f, _ := os.Open(fp)

		buf := make([]byte, 32*1024) // define your buffer size here.

		for {
			n, err := f.Read(buf)

			if n > 0 {
				for _, k := range strings.Split(string(buf[:n]), "\n") {
					if k == "" {
						continue
					}
					k = strings.ToLower(k)
					if !stringExist(k, keywords) {
						keywords = append(keywords, k)
						kwmap[k] = 0
					}
				}
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("read %d bytes: %v", n, err)
				break
			}
		}

	}
	sort.Strings(keywords)

	return nil
}
