package router

import (
	"fmt"
	"net/http"
)

func postAddressHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

//
func getAllAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "You've requested something\n")
}

func getAddressHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

func putAddressHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

func deleteAddressHandler(w http.ResponseWriter, r *http.Request) {
	//...
}
