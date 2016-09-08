package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index -
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Fprint(w, "Welcome!\n")
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(fmt.Fprint(w, "Welcome!\n")))
	fmt.Fprint(w, "Welcome!\n")
}

// TestGET -
func TestGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Test=GET!\n")
}

// TestPOST -
func TestPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Test=POST!\n")
}

// Hello -
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	http.DefaultServeMux = http.NewServeMux()

	if !doesRouteAlreadyExist("GET", "/test", router) {
		fmt.Println("GET 1 does not exist")
		router.Handle("GET", "/test", TestGET)
	}

	if !doesRouteAlreadyExist("GET", "/test", router) {
		fmt.Println("GET 2 does not exist")
		router.Handle("GET", "/test", TestGET)
	}

	h2, p2, e2 := router.Lookup("GET", "/test/1")
	fmt.Println("Exist ", e2)

	if h2 == nil {
		fmt.Println("h2 was nil")
	} else {
		fmt.Println("h2 was not nil")
	}

	if p2 == nil {
		fmt.Println("p2 was nil")
	} else {
		fmt.Println("p2 was not nil")
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

func doesRouteAlreadyExist(method, path string, router *httprouter.Router) bool {
	h, _, _ := router.Lookup(method, path)
	return h != nil
}
