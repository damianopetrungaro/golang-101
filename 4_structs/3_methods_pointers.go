package main

import "fmt"

type user struct {
	id   int
	name string
}

func new(id int, name string) *user {
	return &user{id, name}
}

func (u *user) rename(name string) {
	u.name = name
}

func main() {
	u := new(1, "Mario")
	fmt.Println(u.name)
	u.rename("Luigi")
	fmt.Println(u.name)
}
