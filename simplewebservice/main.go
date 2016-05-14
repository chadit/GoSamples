package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

// to run from the command line type go run main.go
// Open a browser and navigate to localhost:5555
func main() {
	f, _ := os.OpenFile("alogfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)

	log.Printf("Here is some text to write")

	http.Handle("/test", new(MyHandler))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
		//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// pass nil to use the default mux (Multiplexing)
	log.Fatal(http.ListenAndServe(":5555", nil))

}

// MyHandler example struct
type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 not found -" + http.StatusText(404)))
	}
}
