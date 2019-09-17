package main

import "fmt"

func main() {
	printValue("Mario")
	printValue(10)
	printValue(struct{}{})
	printValue(map[string]string{"Man": "Mario", "Woman": "Peach"})
}

func printValue(v interface{}) {
	fmt.Println(v)
}
