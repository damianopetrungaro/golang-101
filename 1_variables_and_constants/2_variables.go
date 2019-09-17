package main

import "regexp"

var globalVariable = "a global string"
var aGlobalRegex = regexp.MustCompile("...")

func main() {
	var name string = "Golang"
	var version float64 = 1.13
	var major int = 1
	// more code here ...
}

// With type inference
func main() {
	var name = "Golang"
	var version = 1.13
	var major = 1
	// more code here ...
}

// With short declaration
func main() {
	name := "Golang"
	version := 1.13
	major := 1
	// more code here ...
}

// With zero-value
func main() {
	var name string     // ""
	var version float64 // 0.0
	var major int       // 0
	// more code here ...
}
