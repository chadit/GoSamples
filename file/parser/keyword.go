package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Keys slice of keys to utlize in parse
type Keys []*Key

// Key information about the key
type Key struct {
	Name  string
	Count int64
}

const keywordPath = "./keywords"

func (h *handler) LoadKeywords() error {
	files, err := getFiles(keywordPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		fp := keywordPath + "/" + file.Name()
		f, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer f.Close()

		buf := make([]byte, 32*2024)

		for {
			n, err := f.Read(buf)

			if n > 0 {
				for _, k := range strings.Split(string(buf[:n]), "\n") {
					if k == "" {
						continue
					}

					if !h.KeyExist(strings.ToLower(k)) {
						h.Keys = append(h.Keys, &Key{Name: k, Count: 0})
					}
				}
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("loadingKeywords : read %d bytes: %v", n, err)
			}
		}

	}
	sort.Sort(h.Keys)
	return nil
}

func (h *handler) KeywordParser(l string) {
	for i := range h.Keys {
		if strings.Contains(l, strings.ToLower(h.Keys[i].Name)) {
			h.Keys[i].Count++
		}
	}
}

func (h *handler) KeyExist(k string) bool {
	for i := range h.Keys {
		if strings.ToLower(h.Keys[i].Name) == k {
			return true
		}
	}
	return false
}

func (k Keys) Len() int           { return len(k) }
func (k Keys) Less(i, j int) bool { return k[i].Name < k[j].Name }
func (k Keys) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
