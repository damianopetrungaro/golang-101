package main

import "errors"

var ErrUserNotFound = errors.New("main: user not found")

func main() {
	var errUserNotFound = errors.New("main: user not found")
	errUserNotLogged := errors.New("main: user not logged")
}
