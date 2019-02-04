package router

import (
	"github.com/gorilla/mux"
)

func CreateRouterHandler() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/address/", postAddressHandler).Methods("POST")
	r.HandleFunc("/address", postAddressHandler).Methods("POST")

	r.HandleFunc("/address/", getAllAddressHandler).Methods("GET")
	r.HandleFunc("/address", getAllAddressHandler).Methods("GET")

	r.HandleFunc("/address/csv/", postCsvAddressHandler).Methods("POST")
	r.HandleFunc("/address/csv", postCsvAddressHandler).Methods("POST")

	r.HandleFunc("/address/csv/", getCsvAddressHandler).Methods("GET")
	r.HandleFunc("/address/csv", getCsvAddressHandler).Methods("GET")

	r.HandleFunc("/address/{addressid}/", getAddressHandler).Methods("GET")
	r.HandleFunc("/address/{addressid}", getAddressHandler).Methods("GET")

	r.HandleFunc("/address/{addressid}/", putAddressHandler).Methods("PUT")
	r.HandleFunc("/address/{addressid}", putAddressHandler).Methods("PUT")

	r.HandleFunc("/address/{addressid}/", deleteAddressHandler).Methods("DELETE")
	r.HandleFunc("/address/{addressid}", deleteAddressHandler).Methods("DELETE")

	return r
}
