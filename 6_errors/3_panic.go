package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	os.Exit(0)
	i, err := strconv.Atoi("Mario")
	switch err := errors.Cause(err).(type) {
	case *strconv.NumError:
		fmt.Println(err)
	default:
		fmt.Println(i)
	}
}
