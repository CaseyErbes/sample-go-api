package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestAddressDataFunc1 tests the capability to create, get all, get individual, update, and delete addresses
func TestAddressDataFunc1(t *testing.T) {
	dbCleanup := InitTestDb()
	defer dbCleanup()

	// create the new address
	firstName, lastName, email, phoneNumber := "Mister", "Test", "mistertest@example.com", "123-456-7890"
	addressId, err := CreateAddress(firstName, lastName, email, phoneNumber)
	assert.Nil(t, err)

	// count the new number of addresses
	addressSlice, err := GetAllAddresses()
	assert.Nil(t, err)
	assert.True(t, len(addressSlice) == 1)

	// get the newly created address and check its properties
	address, err := GetAddress(addressId)
	assert.Nil(t, err)
	assert.True(t, address.AddressId == addressId)
	assert.True(t, address.FirstName == firstName)
	assert.True(t, address.LastName == lastName)
	assert.True(t, address.Email == email)
	assert.True(t, address.PhoneNumber == phoneNumber)

	// update the address
	newFirstName := "Doctor"
	err = UpdateAddress(addressId, newFirstName, lastName, "", "") // demonstrate that empty strings will not update the address entry
	assert.Nil(t, err)

	// check that the address has the updated property
	address, err = GetAddress(addressId)
	assert.Nil(t, err)
	assert.True(t, address.AddressId == addressId)
	assert.True(t, address.FirstName == newFirstName)
	assert.True(t, address.LastName == lastName)
	assert.True(t, address.Email == email)
	assert.True(t, address.PhoneNumber == phoneNumber)

	// delete the address created in this test
	err = DeleteAddress(addressId)
	assert.Nil(t, err)

	// verify that we have the same number of addresses as we started with
	addressSlice, err = GetAllAddresses()
	assert.Nil(t, err)
	assert.True(t, len(addressSlice) == 0)
}

// TestAddressDataFunc2 tests the capability to create a large quantity of addresses,
// get them all, and then delete them all
func TestAddressDataFunc2(t *testing.T) {
	dbCleanup := InitTestDb()
	defer dbCleanup()

	var createdAddressIds []AddressId

	// create 100 new addresses
	newAddressCount := 100
	for i := 1; i <= newAddressCount; i++ {
		firstName, lastName, email, phoneNumber := "Mister", "Test", fmt.Sprintf("mistertest%v@example.com", i), "123-456-7890"
		addressId, err := CreateAddress(firstName, lastName, email, phoneNumber)
		assert.Nil(t, err)
		createdAddressIds = append(createdAddressIds, addressId)
	}

	// count the new number of addresses
	addressSlice, err := GetAllAddresses()
	assert.Nil(t, err)
	assert.True(t, len(addressSlice) == newAddressCount)

	// delete all the created addresses
	for _, addressId := range createdAddressIds {
		err = DeleteAddress(addressId)
		assert.Nil(t, err)
	}

	// verify that we have zero addresses after deleting
	addressSlice, err = GetAllAddresses()
	assert.Nil(t, err)
	assert.True(t, len(addressSlice) == 0)
}
