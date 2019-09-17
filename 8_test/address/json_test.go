package address

import (
	"testing"
)

func TestMarshalJSON(t *testing.T) {

	t.Log("Given an address struct")
	{
		a := Address{
			firstName:   "Damiano",
			lastName:    "Petrungaro",
			address:     "Insulindeweg 17B",
			city:        "Amsterdam",
			postcode:    "1094 NX",
			state:       "Netherlands",
			phoneNumber: "+393332222111",
		}

		t.Log("\tThen an address json MUST be retrieved")
		{
			byte, err := a.MarshalJSON()
			if err != nil {
				t.Fatalf("\t\tAn error occurred marshalling the address. Details: %s", err)
			}

			t.Log("\tAnd the json MUST match the fields of the struct")
			{
				json := `{"first_name":"Damiano","last_name":"Petrungaro","address":"Insulindeweg 17B","city":"Amsterdam","postcode":"1094 NX","state":"Netherlands","phone_number":"+393332222111"}`
				if string(byte) != json {
					t.Error("\t\tThe json are different")
					t.Errorf("\t\tWant: %s", json)
					t.Errorf("\t\tGot:  %s", string(byte))
				}
			}
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {

	t.Log("Given a valid address json")
	{
		got := `{"first_name":"Damiano","last_name":"Petrungaro","address":"Insulindeweg 17B","city":"Amsterdam","postcode":"1094 NX","state":"Netherlands","phone_number":"+393332222111"}`

		t.Log("\tThen an address struct MUST be created")
		{
			a := Address{}
			err := a.UnmarshalJSON([]byte(got))
			if err != nil {
				t.Fatalf("\t\tAn error occurred unmarshalling the address. Details: %s", err)
			}

			t.Log("\tAnd the address struct fields MUST match the firstName used in the JSON")
			{
				want := "Damiano"
				if a.firstName != want {
					t.Error("\t\tThe first name is different")
					t.Errorf("\t\tWant: %s", want)
					t.Errorf("\t\tGot:  %s", a.firstName)
				}
			}
			t.Log("\tAnd the address struct fields MUST match the lastName used in the JSON")
			{
				want := "Petrungaro"
				if a.lastName != want {
					t.Error("\t\tThe last name is different")
					t.Errorf("\t\tWant: %s", want)
					t.Errorf("\t\tGot:  %s", a.lastName)
				}
			}
			t.Log("\tAnd the address struct fields MUST match the address used in the JSON")
			{
				want := "Insulindeweg 17B"
				if a.address != want {
					t.Error("\t\tThe address is different")
					t.Errorf("\t\tWant: %s", want)
					t.Errorf("\t\tGot:  %s", a.address)
				}
			}
			t.Log("\tAnd the address struct fields MUST match the city used in the JSON")
			{
				want := "Amsterdam"
				if a.city != want {
					t.Error("\t\tThe city is different")
					t.Errorf("\t\tWant: %s", want)
					t.Errorf("\t\tGot:  %s", a.city)
				}
			}
			t.Log("\tAnd the address struct fields MUST match the postcode used in the JSON")
			{
				want := "1094 NX"
				if a.postcode != want {
					t.Error("\t\tThe postcode is different")
					t.Errorf("\t\tWant: %s", want)
					t.Errorf("\t\tGot:  %s", a.postcode)
				}
			}
			t.Log("\tAnd the address struct fields MUST match the state used in the JSON")
			{
				want := "Netherlands"
				if a.state != want {
					t.Error("\t\tThe state is different")
					t.Errorf("\t\tWant: %s", want)
					t.Errorf("\t\tGot:  %s", a.state)
				}
			}
			t.Log("\tAnd the address struct fields MUST match the phoneNumber used in the JSON")
			{
				want := "+393332222111"
				if a.phoneNumber != want {
					t.Error("\t\tThe phone number is different")
					t.Errorf("\t\tWant: %s", want)
					t.Errorf("\t\tGot:  %s", a.phoneNumber)
				}
			}
		}
	}
}
