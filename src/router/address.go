package router

import (
	"data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// POST /address
func postAddressHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Warning - ReadAll() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var address data.Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Printf("Warning - Unmarshal() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	addressId, err := data.CreateAddress(address.FirstName, address.LastName, address.Email, address.PhoneNumber)
	if err != nil {
		fmt.Printf("Warning - CreateAddress() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("addressid", string(addressId))
	w.WriteHeader(201)
}

// GET /address
func getAllAddressHandler(w http.ResponseWriter, r *http.Request) {
	addressSlice, err := data.GetAllAddresses()
	if err != nil {
		fmt.Printf("Warning - GetAllAddresses() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(addressSlice)
	if err != nil {
		fmt.Printf("Warning - encoder.Encode() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// GET /address/{addressid}
func getAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId := data.AddressId(vars["addressid"])
	address, err := data.GetAddress(addressId)
	if err != nil {
		fmt.Printf("Warning - GetAddress() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else if address == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(address)
	if err != nil {
		fmt.Printf("Warning - encoder.Encode() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// PUT /address/{addressid}
func putAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId := data.AddressId(vars["addressid"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Warning - ReadAll() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var address data.Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Printf("Warning - Unmarshal() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = data.UpdateAddress(addressId, address.FirstName, address.LastName, address.Email, address.PhoneNumber)
	if err != nil {
		fmt.Printf("Warning - UpdateAddress() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.WriteHeader(204)
}

// DELETE /address/{addressid}
func deleteAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId := data.AddressId(vars["addressid"])
	err := data.DeleteAddress(addressId)
	if err != nil {
		fmt.Printf("Warning - DeleteAddress() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.WriteHeader(204)
}
