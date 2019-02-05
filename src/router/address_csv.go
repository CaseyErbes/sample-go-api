package router

import (
	"data"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// POST /address/csv
func postCsvAddressHandler(w http.ResponseWriter, r *http.Request) {
	csvReader := csv.NewReader(r.Body)
	var csvData [][]string
	for {
		entry, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Warning - csv.Reader.Read() failed with error '%v'\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		csvData = append(csvData, entry)
	}

	colIdxMap := map[string]int{
		"firstName":   -1,
		"lastName":    -1,
		"email":       -1,
		"phoneNumber": -1,
	}

	for i, csvRow := range csvData {
		// header should be first entry
		if i == 0 {
			for j, colName := range csvRow {
				switch {
				case strings.EqualFold("firstName", colName):
					colIdxMap["firstName"] = j
				case strings.EqualFold("lastName", colName):
					colIdxMap["lastName"] = j
				case strings.EqualFold("email", colName):
					colIdxMap["email"] = j
				case strings.EqualFold("phoneNumber", colName):
					colIdxMap["phoneNumber"] = j
				}
			}
			for k, v := range colIdxMap {
				if v == -1 {
					fmt.Printf("Warning - colName '%s' not found in CSV\n", k)
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					return
				}
			}
		} else {
			_, err := data.CreateAddress(
				csvRow[colIdxMap["firstName"]],
				csvRow[colIdxMap["lastName"]],
				csvRow[colIdxMap["email"]],
				csvRow[colIdxMap["phoneNumber"]],
			)
			if err != nil {
				fmt.Printf("Warning - CreateAddress() failed with error '%v'\n", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
	}
	w.WriteHeader(201)
}

// GET /address/csv
func getCsvAddressHandler(w http.ResponseWriter, r *http.Request) {
	addressSlice, err := data.GetAllAddresses()
	if err != nil {
		fmt.Printf("Warning - GetAllAddresses() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	disposition := fmt.Sprintf("filename=\"%s\"", "address_listing.csv")
	w.Header().Set("Content-Disposition", "attachment; "+disposition)
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(200)
	csvWriter := csv.NewWriter(w)
	addressData := [][]string{{"addressId", "firstName", "lastName", "email", "phoneNumber"}}
	for _, address := range addressSlice {
		addressData = append(addressData, []string{string(address.AddressId), address.FirstName, address.LastName, address.Email, address.PhoneNumber})
	}

	csvWriter.WriteAll(addressData)
	if err = csvWriter.Error(); err != nil {
		fmt.Printf("Warning - WriteAll() failed with error '%v'\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
