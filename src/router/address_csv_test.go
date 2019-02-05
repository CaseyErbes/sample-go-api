package router

import (
	"data"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
)

// function gets csv listing all address entries
func getAddressCsv(t *testing.T, baseUrl string) []byte {
	url := baseUrl + "/address/csv"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	require.Nil(t, err)
	res, err := client.Do(req)
	require.Nil(t, err)
	require.True(t, res.StatusCode == 200)
	csvBytes, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	return csvBytes
}

// function uploads a csv of address data
func postAddressCsv(t *testing.T, baseUrl string, fileName string) {
	csv, err := os.Open(fileName)
	require.Nil(t, err)
	defer csv.Close()
	url := baseUrl + "/address/csv"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, csv)
	require.Nil(t, err)
	req.Header.Set("content-type", "text/csv")
	res, err := client.Do(req)
	require.Nil(t, err)
	require.True(t, res.StatusCode == 201)
}

// TestGetAddressCsv tests generating a csv listing all address entries
func TestGetAddressCsv(t *testing.T) {
	dbCleanup := data.InitTestDb()
	defer dbCleanup()

	ts := httptest.NewServer(CreateRouterHandler())
	defer ts.Close()

	// create 100 new addresses
	newAddressCount := 100
	for i := 1; i <= 100; i++ {
		firstName, lastName, email, phoneNumber := "Mister", "Test", fmt.Sprintf("mistertest%v@example.com", i), "123-456-7890"
		addressIdStr := createAddress(t, ts.URL, firstName, lastName, email, phoneNumber)
		assert.True(t, len(addressIdStr) > 0)
	}

	// get generated csv and check that it is the expected length
	csvBytes := getAddressCsv(t, ts.URL)
	re := regexp.MustCompile("\n")
	newLines := re.FindAllStringIndex(string(csvBytes[:]), -1)
	assert.Equal(t, len(newLines)-1, newAddressCount) // -1 because one of the lines is the csv header

	// output csv file for analysis
	err := ioutil.WriteFile("/tmp/csv-test-output.csv", csvBytes, 0644)
	assert.Nil(t, err)
}

// TestPostAddressCsv1 tests uploading a csv listing all address entries into the database
func TestPostAddressCsv1(t *testing.T) {
	dbCleanup := data.InitTestDb()
	defer dbCleanup()

	ts := httptest.NewServer(CreateRouterHandler())
	defer ts.Close()

	// post csv file containing address entries
	postAddressCsv(t, ts.URL, "test1.csv") // file contains 100 entries

	// count the new number of addresses
	addressSlice := getAllAddresses(t, ts.URL)
	assert.True(t, len(addressSlice) == 100)
}

// TestPostAddressCsv2 tests uploading a csv listing all address entries into the database,
// but this time the lastName col comes before the firstName col in the csv
func TestPostAddressCsv2(t *testing.T) {
	dbCleanup := data.InitTestDb()
	defer dbCleanup()

	ts := httptest.NewServer(CreateRouterHandler())
	defer ts.Close()

	// post csv file containing address entries
	postAddressCsv(t, ts.URL, "test2.csv") // file contains 100 entries

	// count the new number of addresses
	addressSlice := getAllAddresses(t, ts.URL)
	assert.True(t, len(addressSlice) == 100)
}
