// package main - read json, cleanup json swagger file, convert to yaml
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	f, err := os.OpenFile("./availability_v3.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	li := fetchLines(f)
	if len(li) > 0 {
		li = li[:len(li)-3]
		fli := li[0]
		li = li[9:]
		li = append([]string{fli}, li...)
	}

	filePath := "./availability_v31.json"
	writeLines(li, filePath)

	ft, _ := os.Open(filePath)
	data, _ := ioutil.ReadAll(ft)
	ft.Close()
	os.Remove(filePath)

	var obj interface{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Printf("Error unmarshalling input. Is it valid JSON? : %v\n", err)
		os.Exit(-1)
	}

	if obj != nil {
		yamlBytes, err := yaml.Marshal(obj)
		if err != nil {
			fmt.Printf("Error marshaling into YAML from JSON file : %v\n", err)
			os.Exit(-1)
		}
		writeFile(yamlBytes, "./availability.yml")
	}

}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

/*
func writeFile(data, file string) {
	f, fError := os.Create(file)
	if fError != nil {
		fmt.Println(fError)
	}
	defer f.Close()
	f.WriteString(data)
}
*/

func writeFile(data []byte, file string) {
	f, fError := os.Create(file)
	if fError != nil {
		fmt.Println(fError)
	}
	defer f.Close()
	f.Write(data)
}

func fetchLines(f *os.File) []string {
	var li []string

	r := bufio.NewReader(f)
	s, e := readln(r)
	li = addLine(s, li)
	for e == nil {
		s, e = readln(r)
		li = addLine(s, li)
	}
	return li
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func addLine(s string, li []string) []string {
	if s != "" {
		li = append(li, s)
	}
	return li
}
