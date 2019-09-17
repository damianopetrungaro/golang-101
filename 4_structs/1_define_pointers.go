package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate *birthDate
}

type birthDate struct {
	t time.Time
}

func main() {
	u := user{
		firstName: "Mario",
		lastName:  "Super",
		birthDate: &birthDate{
			time.Now(),
		},
	}

	fmt.Println(u.firstName, u.lastName, u.birthDate.t.Unix())
}

// Without keys
func main() {
	u := user{
		"Mario",
		"Super",
		&birthDate{
			time.Now(),
		},
	}

	fmt.Println(u.firstName, u.lastName, u.birthDate.t.Unix())
}

// With zero values
func main() {
	u := user{}

	fmt.Println(u.firstName, u.lastName, u.birthDate.t.Unix())
}
