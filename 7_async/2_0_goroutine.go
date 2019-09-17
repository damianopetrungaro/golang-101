package main

import (
	"fmt"
	"time"
)

func main() {
	for _, v := range [5]string{"A", "B", "C", "D", "E"} {
		go fmt.Println(v)
	}
	time.Sleep(3 * time.Second)
}
