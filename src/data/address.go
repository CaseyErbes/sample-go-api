package data

import (
	"errors"
	"github.com/google/uuid"
)

type AddressId string

type Address struct {
	AddressId   AddressId `json:addressId`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
}

var (
	addressInvalidEmailError       = errors.New("address:bad_email")
	addressInvalidPhoneNumberError = errors.New("address:bad_phone_number")
	addressCreateInsertError       = errors.New("address:create_insert_error")
	addressGetAllQueryError        = errors.New("address:get_all_query_error")
	addressGetAllRowScanError      = errors.New("address:get_all_row_scan_error")
	addressGetAllRowsError         = errors.New("address:get_all_rows_error")
	addressGetQueryError           = errors.New("address:get_query_error")
	addressGetRowScanError         = errors.New("address:get_row_scan_error")
	addressUpdateAddressGetError   = errors.New("address:update_address_get_error")
	addressUpdateStatementError    = errors.New("address:update_statement_error")
	addressDeleteStatementError    = errors.New("address:delete_statement_error")
)

func CreateAddress(firstName string, lastName string, email string, phoneNumber string) (AddressId, error) {
	if len(email) == 0 {
		return AddressId(""), addressInvalidEmailError
	}
	if len(phoneNumber) == 0 {
		return AddressId(""), addressInvalidEmailError
	}
	addressIdStr := uuid.New().String()
	statementStr := `INSERT INTO address (id, firstname, lastname, email, phoneNumber)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(statementStr, addressIdStr, firstName, lastName, email, phoneNumber)
	if err != nil {
		return AddressId(""), addressCreateInsertError
	}
	return AddressId(addressIdStr), nil
}

func GetAllAddresses() ([]Address, error) {
	var addressSlice []Address
	queryStr := `SELECT id, firstname, lastname, email, phoneNumber FROM address`
	rows, err := db.Query(queryStr)
	if err != nil {
		return []Address{}, addressGetAllQueryError
	}
	defer rows.Close()
	for rows.Next() {
		var addressIdStr, firstName, lastName, email, phoneNumber string
		err = rows.Scan(&addressIdStr, &firstName, &lastName, &email, &phoneNumber)
		if err != nil {
			return []Address{}, addressGetAllRowScanError
		}
		address := new(Address)
		address.AddressId = AddressId(addressIdStr)
		address.FirstName = firstName
		address.LastName = lastName
		address.Email = email
		address.PhoneNumber = phoneNumber
		addressSlice = append(addressSlice, *address)
	}
	err = rows.Err()
	if err != nil {
		return []Address{}, addressGetAllRowsError
	}
	return addressSlice, nil
}

func GetAddress(addressId AddressId) (*Address, error) {
	address := new(Address)
	queryStr := `SELECT id, firstname, lastname, email, phoneNumber FROM address
		WHERE id = $1`
	rows, err := db.Query(queryStr, string(addressId))
	if err != nil {
		return new(Address), addressGetQueryError
	}
	defer rows.Close()
	for rows.Next() {
		var addressIdStr, firstName, lastName, email, phoneNumber string
		err = rows.Scan(&addressIdStr, &firstName, &lastName, &email, &phoneNumber)
		if err != nil {
			return new(Address), addressGetRowScanError
		}
		address.AddressId = AddressId(addressIdStr)
		address.FirstName = firstName
		address.LastName = lastName
		address.Email = email
		address.PhoneNumber = phoneNumber
	}
	err = rows.Err()
	if err != nil {
		return new(Address), addressGetAllRowsError
	}
	return address, nil
}

func UpdateAddress(addressId AddressId, firstName string, lastName string, email string, phoneNumber string) error {
	currentAddress, err := GetAddress(addressId)
	if err != nil {
		return addressUpdateAddressGetError
	}
	// if empty, email will not be updated
	if len(email) == 0 {
		email = currentAddress.Email
	}
	// if empty, phonenumber will not be updated
	if len(phoneNumber) == 0 {
		phoneNumber = currentAddress.PhoneNumber
	}
	statementStr := `UPDATE address SET
		firstname = $2, lastname = $3, email = $4, phoneNumber = $5
        	WHERE id = $1`
	_, err = db.Exec(statementStr, string(addressId), firstName, lastName, email, phoneNumber)
	if err != nil {
		return addressUpdateStatementError
	}
	return nil
}

func DeleteAddress(addressId AddressId) error {
	statementStr := `DELETE FROM address
        	WHERE id = $1`
	_, err := db.Exec(statementStr, string(addressId))
	if err != nil {
		return addressDeleteStatementError
	}
	return nil
}
