package main

// A simple if statement
func main() {
	a := 1
	if a > 10 {
		// Do something
	}
	// Continue here
}

// An if/else statement
func main() {
	i := len("Mario")
	if i < 10 {
		// Do something
	} else if i > 10 {
		// Do something
	} else {
		// Do something
	}
}

// A better if/else statement
func main() {
	i := len("Mario")
	if i < 10 {
		// Do something
		return
	}
	if i > 10 {
		// Do something
		return
	}
	// Do something
}
