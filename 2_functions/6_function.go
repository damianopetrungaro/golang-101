package main

import (
	"fmt"
	"os/exec"
)

// A simple function
func funcName() {
	// Do something
}

// With a value
func funcName(i int) {
	// Do something
}

// With three integers
func funcName(a, b, c int) {
	// Do something
}

// With an integer and a string
func funcName(i int, s string) {
	// Do something
}

// With a return type
func funcName(i int, s string) bool {
	return false
}

// With n return types :wat:
func funcName(i int, s string) (string, bool) {
	return "", true
}

// With an error
func funcName(i int, s string) (string, error) {
	return "", exec.ErrNotFound
}

// With named return values
func funcName() (name string, surname string, nickname string) {
	return
}

// With a named return value changed
func funcName() (name string, surname string, nickname string) {
	name = "Mario"
	return
}

// Lambda
func main() {
	a := func(name string) {
		fmt.Println(name)
	}

	a("Mario")
	a("Luigi")

	func(name string) {
		fmt.Println(name)
	}("Peach")
}
