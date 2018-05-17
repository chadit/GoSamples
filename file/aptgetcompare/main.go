package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	desktopFile  = "./desktop.txt"
	notebookFile = "./notebook.txt"
)

func main() {
	desktopLines, err := readFile(desktopFile)
	if err != nil {
		fmt.Println("desktop :", err)
		return
	}

	notebookLines, err := readFile(notebookFile)
	if err != nil {
		fmt.Println("notebook :", err)
		return
	}

	fmt.Printf("\n%v\n%v\n", desktopLines[0], notebookLines[0])

	var missingPackagesOnNotebook []string

	for i := range desktopLines {
		if !packageExistOnBoth(notebookLines, desktopLines[i]) {
			missingPackagesOnNotebook = append(missingPackagesOnNotebook, desktopLines[i])
		}
	}

	fmt.Println(len(missingPackagesOnNotebook))

	fmt.Printf("sudo apt install -y %s", strings.Join(missingPackagesOnNotebook, " "))

}

func packageExistOnBoth(items []string, key string) bool {

	for i := range items {

		if items[i] == key {
			return true
		}
	}

	return false

}

func readFile(file string) ([]string, error) {
	var lines []string

	f, err := os.Open(file)
	if err != nil {
		return lines, err
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

				k = strings.Replace(k, "\t", "", -1)
				k = strings.Replace(k, "install", "", -1)

				lines = append(lines, k)
			}

		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("loadingKeywords : read %d bytes: %v", n, err)
		}
	}

	return lines, nil
}
