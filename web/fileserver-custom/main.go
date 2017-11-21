package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(fmt.Sprintf("%s/public%s", parentFilePathHelper(), r.URL.Path))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		defer f.Close()
		// need to set the content type so that it is rendered correctly, otherwise it renders as plain text
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "png"):
			contentType = "image/png"
		default:
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})

	http.ListenAndServe(":8000", nil)
}

// urlHelper - returns the absolute path for the main file, go run main.go does not have the same path as an executable
func parentFilePathHelper() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return strings.Replace(path.Dir(filename), "src/main", "", 1)
}
