package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Lines []*Line
type Line struct {
	Hash        string
	Text        string
	TokenCounts []int
	DupCount    int64
	Length      int
}

const textPath = "./text_files"

func (h *handler) ProcessTextFiles() error {
	files, err := getFiles(textPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		fp := textPath + "/" + file.Name()
		f, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer f.Close()

		buf := make([]byte, 32*2024)

		for {
			n, err := f.Read(buf)

			if n > 0 {
				for _, k := range strings.Split(string(buf[:n]), "\r\n") {
					if len(k) == 0 {
						continue
					}
					h.ParseLine(k)
				}
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("read %d bytes: %v", n, err)
			}
		}
	}
	return nil
}

func (h *handler) ParseLine(l string) {
	lLen := len(l)
	if lLen == 0 {
		return
	}

	ex, i := h.LineExist(strings.ToLower(l))
	if ex {
		// if Line exist increment count, if zero, increase twice for two lines
		if h.Lines[i].DupCount == 0 {
			h.Lines[i].DupCount++
		}
		h.Lines[i].DupCount++
	} else {
		h.Lines[i].Length = lLen
	}

	if token := len(strings.Fields(l)); token > 0 {
		h.Lines[i].TokenCounts = append(h.Lines[i].TokenCounts, token)
	}

	h.KeywordParser(l)
}

// LineExist checks lines slice to see if line exist, returns if it does and the index
func (h *handler) LineExist(l string) (bool, int) {
	for i := range h.Lines {
		if strings.ToLower(h.Lines[i].Text) == l {
			return true, i
		}
	}
	nl := Line{Text: l}
	h.Lines = append(h.Lines, &nl)

	// return index of newly added line item
	return false, len(h.Lines) - 1
}

func (h *handler) StatsParser() {
	ld := []int{}
	td := []int{}
	for i := range h.Lines {
		ls := h.Lines[i]
		ld = append(ld, ls.Length)
		td = append(td, ls.TokenCounts...)
		h.Stats.LineDupCount += ls.DupCount
	}

	h.Stats.LineMedian = calcMedian(ld)
	h.Stats.LineStDev = calcSD(ld)
	h.Stats.TokenMedian = calcMedian(td)
	h.Stats.TokenStDev = calcSD(td)
}
