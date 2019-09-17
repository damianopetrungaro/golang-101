package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		fmt.Println(s)
	}("A")

	go func(s string) {
		fmt.Println(s)
	}("B")

	time.Sleep(3 * time.Second)
}
