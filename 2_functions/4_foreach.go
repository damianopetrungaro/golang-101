package main

// A simple for loop (again)
func main() {
	list := []int{1, 10, 234, 63, 32, 8, 24, 8}
	for i, v := range list {
		// Do something
	}
}

// A simple for loop (again) ignoring the index
func main() {
	list := []int{1, 10, 234, 63, 32, 8, 24, 8}
	for _, v := range list {
		// Do something
	}
}
