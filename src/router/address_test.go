package router

import (
	"data"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// function creates a new address entry
func createAddress(t *testing.T, baseUrl string, firstName string, lastName string, email string, phoneNumber string) string {
	payload := fmt.Sprintf(`{
                "firstName": "%s",
                "lastName": "%s",
                "email": "%s",
                "phoneNumber": "%s"
        }`, firstName, lastName, email, phoneNumber)
	url := baseUrl + "/address/"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	require.Nil(t, err)
	res, err := client.Do(req)
	require.Nil(t, err)
	require.True(t, res.StatusCode == 201)
	return res.Header.Get("addressid")
}

// function lists all address entries
func getAllAddresses(t *testing.T, baseUrl string) []data.Address {
	url := baseUrl + "/address/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)
	res, err := client.Do(req)
	require.Nil(t, err)
	require.True(t, res.StatusCode == 200)
	result, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	var addressSlice []data.Address
	err = json.Unmarshal(result, &addressSlice)
	require.Nil(t, err)
	return addressSlice
}

// function returns an individual address entry
func getAddress(t *testing.T, baseUrl string, addressIdStr string) *data.Address {
	url := baseUrl + "/address/" + addressIdStr
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	require.Nil(t, err)
	res, err := client.Do(req)
	require.Nil(t, err)
	require.True(t, res.StatusCode == 200)
	result, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	address := new(data.Address)
	err = json.Unmarshal(result, &address)
	require.Nil(t, err)
	return address
}

// function modifies an address entry
func updateAddress(t *testing.T, baseUrl string, addressIdStr string, firstName string, lastName string, email string, phoneNumber string) {
	payload := fmt.Sprintf(`{
                "firstName": "%s",
                "lastName": "%s",
                "email": "%s",
                "phoneNumber": "%s"
        }`, firstName, lastName, email, phoneNumber)
	url := baseUrl + "/address/" + addressIdStr
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	require.Nil(t, err)
	res, err := client.Do(req)
	require.Nil(t, err)
	require.True(t, res.StatusCode == 204)
}

// function deletes an address entry
func deleteAddress(t *testing.T, baseUrl string, addressIdStr string) {
	url := baseUrl + "/address/" + addressIdStr
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	require.Nil(t, err)
	res, err := client.Do(req)
	require.Nil(t, err)
	require.True(t, res.StatusCode == 204)
}

// TestAddressHttpFunc1 tests the capability to create, get all, get individual, update, and delete addresses
func TestAddressHttpFunc1(t *testing.T) {
	dbCleanup := data.InitTestDb()
	defer dbCleanup()

	ts := httptest.NewServer(CreateRouterHandler())
	defer ts.Close()

	// create the new address
	firstName, lastName, email, phoneNumber := "Mister", "Test", "mistertest@example.com", "123-456-7890"
	addressIdStr := createAddress(t, ts.URL, firstName, lastName, email, phoneNumber)
	assert.True(t, len(addressIdStr) > 0)

	// count the new number of addresses
	addressSlice := getAllAddresses(t, ts.URL)
	assert.True(t, len(addressSlice) == 1)

	// get the newly created address and check its properties
	address := getAddress(t, ts.URL, addressIdStr)
	assert.True(t, string(address.AddressId) == addressIdStr)
	assert.True(t, address.FirstName == firstName)
	assert.True(t, address.LastName == lastName)
	assert.True(t, address.Email == email)
	assert.True(t, address.PhoneNumber == phoneNumber)

	// update the address
	newFirstName, newLastName, newEmail, newPhoneNumber := "Doctor", "Test II", "doctortest2@example.com", "098-765-4321"
	updateAddress(t, ts.URL, addressIdStr, newFirstName, newLastName, newEmail, newPhoneNumber)

	// check that the address has the updated property
	address = getAddress(t, ts.URL, addressIdStr)
	assert.True(t, string(address.AddressId) == addressIdStr)
	assert.True(t, address.FirstName == newFirstName)
	assert.True(t, address.LastName == newLastName)
	assert.True(t, address.Email == newEmail)
	assert.True(t, address.PhoneNumber == newPhoneNumber)

	// delete the address created in this test
	deleteAddress(t, ts.URL, addressIdStr)

	// verify that we have zero addresses after deleting
	addressSlice = getAllAddresses(t, ts.URL)
	assert.True(t, len(addressSlice) == 0)
}

// TestAddressHttpFunc2 tests the capability to create a large quantity of addresses,
// get them all, and then delete them all
func TestAddressHttpFunc2(t *testing.T) {
	dbCleanup := data.InitTestDb()
	defer dbCleanup()

	ts := httptest.NewServer(CreateRouterHandler())
	defer ts.Close()

	var createdAddressIdStrs []string

	// create 100 new addresses
	newAddressCount := 100
	for i := 1; i <= 100; i++ {
		firstName, lastName, email, phoneNumber := "Mister", "Test", fmt.Sprintf("mistertest%v@example.com", i), "123-456-7890"
		addressIdStr := createAddress(t, ts.URL, firstName, lastName, email, phoneNumber)
		assert.True(t, len(addressIdStr) > 0)
		createdAddressIdStrs = append(createdAddressIdStrs, addressIdStr)
	}

	// count the new number of addresses
	addressSlice := getAllAddresses(t, ts.URL)
	assert.True(t, len(addressSlice) == newAddressCount)

	// delete the addresses created in this test
	for _, addressIdStr := range createdAddressIdStrs {
		deleteAddress(t, ts.URL, addressIdStr)
	}

	// verify that we have zero addresses after deleting
	addressSlice = getAllAddresses(t, ts.URL)
	assert.True(t, len(addressSlice) == 0)
}
