package address

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Errors used on failure during marshal/unmarshal
var (
	ErrMarshalAddress   = errors.New("address: an error occurred marshalling")
	ErrUnmarshalAddress = errors.New("address: an error occurred unmarshalling")
)

type jsonAddress struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postcode"`
	State       string `json:"state"`
	PhoneNumber string `json:"phone_number"`
}

// MarshalJSON implements json.Marshaler
func (a Address) MarshalJSON() ([]byte, error) {
	if a == (Address{}) {
		return json.Marshal(struct{}{})
	}

	data, err := json.Marshal(jsonAddress{
		FirstName:   a.firstName,
		LastName:    a.lastName,
		Address:     a.address,
		City:        a.city,
		Postcode:    a.postcode,
		State:       a.state,
		PhoneNumber: a.phoneNumber,
	})

	if err != nil {
		return nil, errors.Wrap(ErrMarshalAddress, err.Error())
	}

	return data, nil
}

// UnmarshalJSON implements json.Unmarshaler
func (a *Address) UnmarshalJSON(data []byte) error {

	jsonAddress := jsonAddress{}
	if err := json.Unmarshal(data, &jsonAddress); err != nil {
		return errors.Wrap(ErrUnmarshalAddress, err.Error())
	}

	address, err := New(
		jsonAddress.FirstName,
		jsonAddress.LastName,
		jsonAddress.Address,
		jsonAddress.City,
		jsonAddress.Postcode,
		jsonAddress.State,
		jsonAddress.PhoneNumber,
	)
	if err != nil {
		return errors.Wrap(ErrUnmarshalAddress, err.Error())
	}

	a.firstName = address.FirstName()
	a.lastName = address.LastName()
	a.address = address.Address()
	a.city = address.City()
	a.postcode = address.Postcode()
	a.state = address.State()
	a.phoneNumber = address.PhoneNumber()

	return nil
}
