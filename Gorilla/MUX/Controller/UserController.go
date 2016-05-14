package Controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte(fmt.Sprintf("Hello %s from usercontroller !", name)))
}
