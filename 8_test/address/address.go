package address

import (
	"strings"

	"github.com/pkg/errors"
)

// Errors returned from address fields validation
var (
	ErrInvalidFirstName   = errors.New("first name in address can't be empty")
	ErrInvalidLastName    = errors.New("last name in address can't be empty")
	ErrInvalidAddress     = errors.New("address1 in address can't be empty")
	ErrInvalidCity        = errors.New("city in address can't be empty")
	ErrInvalidPostcode    = errors.New("postcode in address can't be empty")
	ErrInvalidState       = errors.New("state in address can't be empty")
	ErrInvalidPhoneNumber = errors.New("phone number in address can't be empty")
)

// Address Representation of an address
type Address struct {
	firstName   string
	lastName    string
	address     string
	city        string
	postcode    string
	state       string
	phoneNumber string
}

// New Return an Address value
func New(
	firstName,
	lastName,
	address,
	city,
	postcode,
	state,
	phoneNumber string,
) (Address, error) {

	firstName = strings.TrimSpace(firstName)
	if firstName == "" {
		return Address{}, ErrInvalidFirstName
	}

	lastName = strings.TrimSpace(lastName)
	if lastName == "" {
		return Address{}, ErrInvalidLastName
	}

	address = strings.TrimSpace(address)
	if address == "" {
		return Address{}, ErrInvalidAddress
	}

	city = strings.TrimSpace(city)
	if city == "" {
		return Address{}, ErrInvalidCity
	}

	postcode = strings.TrimSpace(postcode)
	if postcode == "" {
		return Address{}, ErrInvalidPostcode
	}

	state = strings.TrimSpace(state)
	if state == "" {
		return Address{}, ErrInvalidState
	}

	phoneNumber = strings.TrimSpace(phoneNumber)
	if phoneNumber == "" {
		return Address{}, ErrInvalidPhoneNumber
	}

	return Address{
		firstName:   firstName,
		lastName:    lastName,
		address:     address,
		city:        city,
		postcode:    postcode,
		state:       state,
		phoneNumber: phoneNumber,
	}, nil
}

//FirstName Retrieve address firstName
func (a Address) FirstName() string {
	return a.firstName
}

//LastName Retrieve address lastName
func (a Address) LastName() string {
	return a.lastName
}

//Address Retrieve address address
func (a Address) Address() string {
	return a.address
}

//City Retrieve address city
func (a Address) City() string {
	return a.city
}

//Postcode Retrieve address postcode
func (a Address) Postcode() string {
	return a.postcode
}

//State Retrieve address state
func (a Address) State() string {
	return a.state
}

//PhoneNumber Retrieve address phoneNumber
func (a Address) PhoneNumber() string {
	return a.phoneNumber
}

//Equals Return true if the two value objects  are identical
func (a Address) Equals(address Address) bool {
	return a.firstName == address.firstName &&
		a.lastName == address.lastName &&
		a.address == address.address &&
		a.city == address.city &&
		a.postcode == address.postcode &&
		a.state == address.state &&
		a.phoneNumber == address.phoneNumber
}
