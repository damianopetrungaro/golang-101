package main

import "fmt"

func main() {
	ch := make(chan string)
	defer close(ch)

	go func() { ch <- "Hello" }()

	msg := <-ch
	fmt.Println(msg)
}
