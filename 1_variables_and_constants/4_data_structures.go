package main

import "fmt"

// As zero-value
func main() {
	var anArray [1]int
	var aSlice []int
	var aMap map[string]int
}

// With composite literal
func main() {
	anArray := [1]int{2}
	aSlice := []string{"value"}
	aMap := map[string]int{"age": 20}
}

// With length and capacity
func main() {
	var aSlice = make([]string, 2, 10)
	var anotherSlice = new([]string)
}

// What's inside?
func main() {
	var aSlice = make([]string, 4, 5)
	var anEmptySlice = make([]string, 0, 10)
	var aPointerToSlice = new([]string)

	fmt.Printf("%#v\n", aSlice)          // []string{"", "", "", ""}
	fmt.Printf("%#v\n", anEmptySlice)    // []string{}
	fmt.Printf("%#v\n", aPointerToSlice) // &[]string(nil)
}
