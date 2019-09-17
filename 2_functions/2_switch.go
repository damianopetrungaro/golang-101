package main

// A simple switch statement
func main() {
	name := "Mario"
	switch name {
	case "Mario":
		// Do something
	case "Luigi":
		// Do something
	default:
		// Do something
	}
}

// A simple switch statement with assignment
func main() {
	var name interface{}
	name = "Mario"
	switch t := name.(type) {
	case string:
		// Do something with t
	case int:
		// Do something with t
	default:
		// Do something with t
	}
}
