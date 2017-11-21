package main

import (
	"fmt"
	"net/http"
	"path"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("%s/public%s", parentFilePathHelper(), r.URL.Path))
	})

	http.ListenAndServe(":8000", nil)
}

// urlHelper - returns the absolute path for the main file, go run main.go does not have the same path as an executable
func parentFilePathHelper() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return strings.Replace(path.Dir(filename), "src/main", "", 1)
}
