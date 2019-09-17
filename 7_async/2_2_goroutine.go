package main

import (
	"fmt"
	"time"
)

func main() {
	fn := func(s string) {
		fmt.Println(s)
	}

	fn("A")
	fn("B")

	time.Sleep(3 * time.Second)
}
