package main

import (
	"errors"
	"fmt"
	"strconv"
)

func stringLength(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("main: an error occurred transforming string to integer")
	}

	return i, nil
}

// One-line idiomatic error check
func main() {
	if _, err := stringLength("Mario"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Job done!")
}

// Keep in mind the scope of the "if"
func main() {
	i, err := stringLength("Mario")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(i)
}
