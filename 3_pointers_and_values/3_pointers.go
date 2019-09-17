package main

import (
	"fmt"
	"net/http"
)

func main() {
	const url = "http://google.com"
	req, err := http.NewRequest(http.MethodGet, "http://google.com", nil)
	if err != nil {
		fmt.Println("an error occurred", err)
		return
	}

	fmt.Println(req.URL.String())
}

func main() {
	const url = "http://google.com"
	req, err := NewRequest(http.MethodGet, "http://google.com")
	if err != nil {
		fmt.Println("an error occurred", err)
		return
	}

	fmt.Println(req.URL.String())
}

func NewRequest(method, url string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, url, nil)
}
