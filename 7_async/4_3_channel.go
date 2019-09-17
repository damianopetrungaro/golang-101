package main

import "fmt"

func main() {

	msgs := make(chan string, 2)
	defer close(msgs)

	go func() {
		msgs <- "one"
		msgs <- "two"
		msgs <- "three!"
	}()

	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
}
