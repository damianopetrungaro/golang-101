package main

import (
	"regexp"
)

const globalConstant = "a global string"
const r = regexp.MustCompile("") // const initializer regexp.MustCompile("") is not a constant

func main() {
	const name string = "Golang"
	const version float64 = 1.13
	const major int = 1
	// more code here ...
}

// With type inference
func main() {
	const name = "Golang"
	const version = 1.13
	const major = 1
	// more code here ...
}
