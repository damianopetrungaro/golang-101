package main

import "fmt"

type username string
type printUsernameFunc func(username) bool

func printUsername(u username) bool {
	// print the username
	return true
}

func main() {
	mario := username("mario")
	printUsername(mario)
	printUsername("luigi")
}

func main() {
	complexfunction(printUsername)
}

func complexfunction(fn printUsernameFunc) {
	mario := username("mario")
	ok := fn(mario)
	fmt.Println(ok)
	ok = fn("luigi")
	fmt.Println(ok)
}
