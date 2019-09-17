package main

import (
	"fmt"
	"net/url"
)

func main() {
	var s = new(string)
	*s = "another value"
	print(s)

	var u = &url.URL{
		//..
	}
	printUri(u)
}

func print(s *string) {
	fmt.Println(s)
}

func printUri(u *url.URL) {
	fmt.Println(u.String())
}
