package main

import "fmt"

func main() {
	for _, v := range [5]string{"A", "B", "C", "D", "E"} {
		go fmt.Println(v)
	}
}
