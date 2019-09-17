package address

import (
	"fmt"
	"testing"
)

func TestFactoryMethodFailure(t *testing.T) {
	list := []struct {
		got  [7]string
		want error
	}{
		{
			[7]string{" ", "Petrungaro", "A random street", "Amsterdam", "1094 NX", "Netherlands", "+393332222111"},
			ErrInvalidFirstName,
		},
		{
			[7]string{"Damiano", " ", "A random street", "Amsterdam", "1094 NX", "Netherlands", "+393332222111"},
			ErrInvalidLastName,
		},
		{
			[7]string{"Damiano", "Petrungaro", " ", "Amsterdam", "1094 NX", "Netherlands", "+393332222111"},
			ErrInvalidAddress,
		},
		{
			[7]string{"Damiano", "Petrungaro", "A random street", " ", "1094 NX", "Netherlands", "+393332222111"},
			ErrInvalidCity,
		},
		{
			[7]string{"Damiano", "Petrungaro", "A random street", "Amsterdam", " ", "Netherlands", "+393332222111"},
			ErrInvalidPostcode,
		},
		{
			[7]string{"Damiano", "Petrungaro", "A random street", "Amsterdam", "1094 NX", " ", "+393332222111"},
			ErrInvalidState,
		},
		{
			[7]string{"Damiano", "Petrungaro", "A random street", "Amsterdam", "1094 NX", "Netherlands", " "},
			ErrInvalidPhoneNumber,
		},
	}

	for i, d := range list {
		t.Run(fmt.Sprintf("Tets with index %d", i), func(t *testing.T) {
			t.Log("\tGiven an invalid payload")
			{
				t.Log("\tThen the a struct MUST NOT be a valid struct")
				a, err := New(d.got[0], d.got[1], d.got[2], d.got[3], d.got[4], d.got[5], d.got[6])
				if a != (Address{}) {
					t.Errorf("\t\tThe a MUST be an empty struct")
					t.Errorf("\t\tGot: %+v", a)
					t.Errorf("\t\tWant: %+v", Address{})
				}
				t.Log("\tAnd the error MUST match the wanted one")
				if err != d.want {
					t.Errorf("\t\tThe error is not the expected one")
					t.Errorf("\t\tGot: %s", err)
					t.Errorf("\t\tWant: %s", d.want)
				}
			}
		})
	}
}

func TestFactoryMethodSuccess(t *testing.T) {
	list := []struct {
		got  [7]string
		want [7]string
	}{
		{
			[7]string{"   Damiano", "   Petrungaro", "   A random street", "Amsterdam   ", "1094 NX   ", "Netherlands   ", " +39 3926562754"},
			[7]string{"Damiano", "Petrungaro", "A random street", "Amsterdam", "1094 NX", "Netherlands", "+39 3926562754"},
		},
	}

	for i, d := range list {
		t.Run(fmt.Sprintf("Tets with index %d", i), func(t *testing.T) {
			t.Log("\tGiven a payload")
			{
				t.Log("\tThen the struct MUST be populated with no error")
				a, err := New(d.got[0], d.got[1], d.got[2], d.got[3], d.got[4], d.got[5], d.got[6])
				if err != nil {
					t.Errorf("\t\tAn error occurred creating the struct")
					t.Errorf("\t\tGot: %s", err)
				}
				t.Log("\tAnd the struct fields MUST match the firstName used in the payload")
				{
					if a.FirstName() != d.want[0] {
						t.Error("\t\tThe first name is different")
						t.Errorf("\t\tWant: %s", d.want[0])
						t.Errorf("\t\tGot:  %s", a.FirstName())
					}
				}
				t.Log("\tAnd the struct fields MUST match the lastName used in the payload")
				{
					if a.LastName() != d.want[1] {
						t.Error("\t\tThe last name is different")
						t.Errorf("\t\tWant: %s", d.want[1])
						t.Errorf("\t\tGot:  %s", a.LastName())
					}
				}
				t.Log("\tAnd the struct fields MUST match the address used in the payload")
				{
					if a.Address() != d.want[2] {
						t.Error("\t\tThe address is different")
						t.Errorf("\t\tWant: %s", d.want[2])
						t.Errorf("\t\tGot:  %s", a.Address())
					}
				}
				t.Log("\tAnd the struct fields MUST match the city used in the payload")
				{
					if a.City() != d.want[3] {
						t.Error("\t\tThe city is different")
						t.Errorf("\t\tWant: %s", d.want[3])
						t.Errorf("\t\tGot:  %s", a.City())
					}
				}
				t.Log("\tAnd the struct fields MUST match the postcode used in the payload")
				{
					if a.Postcode() != d.want[4] {
						t.Error("\t\tThe postcode is different")
						t.Errorf("\t\tWant: %s", d.want[4])
						t.Errorf("\t\tGot:  %s", a.Postcode())
					}
				}
				t.Log("\tAnd the struct fields MUST match the state used in the payload")
				{
					if a.State() != d.want[5] {
						t.Error("\t\tThe state is different")
						t.Errorf("\t\tWant: %s", d.want[5])
						t.Errorf("\t\tGot:  %s", a.State())
					}
				}
				t.Log("\tAnd the struct fields MUST match the phoneNumber used in the payload")
				{
					if a.PhoneNumber() != d.want[6] {
						t.Error("\t\tThe phone number is different")
						t.Errorf("\t\tWant: %s", d.want[6])
						t.Errorf("\t\tGot:  %s", a.PhoneNumber())
					}
				}
			}
		})
	}
}

func TestEquals(t *testing.T) {
	list := []struct {
		got  [2][7]string
		want bool
	}{
		{
			[2][7]string{
				{"Damiano", "Petrungaro", "A random street", "Amsterdam", "1094 NX", "Netherlands", "+39 3926562754"},
				{"Damiano", "Petrungaro", "A random street", "Amsterdam", "1094 NX", "Netherlands", "+39 3926562754"},
			}, true,
		},
		{
			[2][7]string{
				{"Damiano", "Petrungaro", "A random street", "Amsterdam", "1094 NX", "Netherlands", "+39 3926562754"},
				{"Antonio", "Petrungaro", "A random street", "Amsterdam", "1094 NX", "Netherlands", "+39 3926562754"},
			}, false,
		},
	}

	for i, d := range list {
		t.Run(fmt.Sprintf("Tets with index %d", i), func(t *testing.T) {
			t.Log("\tGiven two payload")
			{

				t.Log("\tThen two structs are created")
				a1 := Address{d.got[0][0], d.got[0][1], d.got[0][2], d.got[0][3], d.got[0][4], d.got[0][5], d.got[0][6]}
				a2 := Address{d.got[1][0], d.got[1][1], d.got[1][2], d.got[1][3], d.got[1][4], d.got[1][5], d.got[1][6]}
				t.Log("\tAnd their equation match the expected one")
				eq := a1.Equals(a2)
				if d.want != eq {
					t.Error("\t\tThe equation do not match")
					t.Errorf("\t\tGot: %t", eq)
					t.Errorf("\t\tWant: %t", d.want)
				}
			}
		})
	}
}
