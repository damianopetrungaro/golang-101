package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "one"
	messages <- "three"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
