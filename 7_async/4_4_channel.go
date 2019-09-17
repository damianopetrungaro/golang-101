package main

import "fmt"

func main() {

	msgs := make(chan string, 3)

	msgs <- "one"
	msgs <- "two"
	msgs <- "three!"
	close(msgs)

	for v := range msgs {
		fmt.Println(v)
	}
}
