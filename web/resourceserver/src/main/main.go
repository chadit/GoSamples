package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
	"strings"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })
	mux := http.NewServeMux()
	mh := &MyHandler{}
	mux.Handle("/", mh)
	//	mh := &MyHandler{}
	//http.Handle("/", mh)

	fmt.Println(http.ListenAndServe(":8000", mux))
}

// MyHandler -
type MyHandler struct {
	http.Handler
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

// ServerHTTP -
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req != nil && req.URL != nil {
		path := fmt.Sprintf("%s/public%s", parentFilePathHelper(), req.URL.Path)
		data, err := ioutil.ReadFile(string(path))
		if err == nil {
			w.Write(data)
		} else {
			fmt.Println("err : ", err)
			w.WriteHeader(404)
			w.Write([]byte("404 - " + http.StatusText(404)))
		}
	} else {
		w.Write([]byte("Hello World"))
	}
}
