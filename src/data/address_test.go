package data

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestAddressFunc1 tests the capability to create, get all, get individual, update, and delete addresses
func TestAddressFunc1(t *testing.T) {
	InitDb()
	// count the number addresses to start out with
	addressSlice, err := GetAllAddresses()
        assert.Nil(t, err)
	initialAddressCount := len(addressSlice)
	// create the new address
	firstName, lastName, email, phoneNumber := "Mister", "Test", "mistertest@example.com", "123-456-7890"
	addressId, err := CreateAddress(firstName, lastName, email, phoneNumber)
	assert.Nil(t, err)
	// count the new number of addresses
	addressSlice, err = GetAllAddresses()
	assert.Nil(t, err)
	assert.True(t, len(addressSlice) == initialAddressCount+1)
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
	err = UpdateAddress(addressId, newFirstName, lastName, "", "")
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
}

// TestAddressFunc2 tests the capability to create a large quantity of addresses,
// get them all, and then delete them all
func TestAddressFunc2(t *testing.T) {
        InitDb()
	var createdAddressIds []AddressId
        // count the number addresses to start out with
        addressSlice, err := GetAllAddresses()
        assert.Nil(t, err)
        initialAddressCount := len(addressSlice)
        // create 5 new address
        firstName, lastName, email, phoneNumber := "Mister", "Test", "mistertest1@example.com", "123-456-7890"
        addressId, err := CreateAddress(firstName, lastName, email, phoneNumber)
        assert.Nil(t, err)
	createdAddressIds = append(createdAddressIds, addressId)
        email = "mistertest2@example.com"
        addressId, err = CreateAddress(firstName, lastName, email, phoneNumber)
        assert.Nil(t, err)
	createdAddressIds = append(createdAddressIds, addressId)
        email = "mistertest3@example.com"
        addressId, err = CreateAddress(firstName, lastName, email, phoneNumber)
        assert.Nil(t, err)
	createdAddressIds = append(createdAddressIds, addressId)
        email = "mistertest4@example.com"
        addressId, err = CreateAddress(firstName, lastName, email, phoneNumber)
        assert.Nil(t, err)
	createdAddressIds = append(createdAddressIds, addressId)
        email = "mistertest5@example.com"
        addressId, err = CreateAddress(firstName, lastName, email, phoneNumber)
        assert.Nil(t, err)
	createdAddressIds = append(createdAddressIds, addressId)
        // count the new number of addresses
        addressSlice, err = GetAllAddresses()
        assert.Nil(t, err)
        assert.True(t, len(addressSlice) == initialAddressCount+5)
	// delete all the created addresses
	for _, addressId := range createdAddressIds {
		err = DeleteAddress(addressId)
		assert.Nil(t, err)
	}
}
