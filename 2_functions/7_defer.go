package main

import "fmt"

func main() {
	defer printZero()
	printOne()
	printTwo()
}

func printZero() {
	fmt.Println(0)
}

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}
