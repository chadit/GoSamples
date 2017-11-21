package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	http.Handler
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%s world", mh.greeting)))
}

func main() {
	http.Handle("/hi", &myHandler{greeting: "Hello"})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Println(http.ListenAndServe(":8000", nil))
}
