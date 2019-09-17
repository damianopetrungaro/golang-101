package main

import "fmt"

type user struct {
	name string
}

type admin struct {
	user
	role string
}

func main() {
	u := user{"Mario"}
	a := admin{user{"Luigi"}, "CTO"}

	fmt.Println(u.name)
	fmt.Println(a.name, a.role)
}
