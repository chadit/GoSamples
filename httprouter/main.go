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
	//	router.GET("/", Index)
	router.Handle("GET", "/test", TestGET)
	router.Handle("POST", "/test", TestPOST)

	//router.GET("/test", TestGET)
	//router.POST("/test", TestPOST)
	//	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func (t T) setupRoute(m, u string, h func(T) T) {
	t.method = m
	t.handler = h
	t.requireAuth = true
	t.Router.OPTIONS(u, t.handleOPTIONS)
	t.Router.Handle(m, u, t.vetHTTPRequest)
}

// T test
type T struct {
	method      string             // indicates HTTP method (e.g: "GET", "POST")
	Router      *httprouter.Router // handles routing concerns
	handler     func(T) T          // handles concerns unique to a particular route
	requireAuth bool               // indicates that route requires authentication
}

// handleOPTIONS handles OPTIONS requests.
func (t T) handleOPTIONS(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if o := r.Header.Get("Origin"); o != "" {
		h := w.Header()
		h.Set("Access-Control-Allow-Origin", o)
		h.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		h.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Referer")
	}

	if r.Method == "OPTIONS" {
		return
	}
}

// vetHTTPRequest is used for vetting HTTP requests.
// Until user is authenticated and authorized, no request processing occurs.
func (t T) vetHTTPRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t.handleOPTIONS(w, r, ps)
}
