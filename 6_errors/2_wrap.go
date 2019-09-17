package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// Wrapping errors
func stringLength(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Wrap(err, "main: an error occurred transforming string to integer")
	}

	return i, nil
}

func main() {
	i, err := stringLength("Mario")
	switch err := errors.Cause(err).(type) {
	case *strconv.NumError:
		fmt.Println(err)
	default:
		fmt.Println(i)
	}
}
