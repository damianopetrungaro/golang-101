package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ch := make(chan int)
	ctx := context.Background()
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	go doAsyncRequest(ch, req, client)

	ctx, _ = context.WithTimeout(ctx, 650*time.Millisecond)

	select {
	case status := <-ch:
		fmt.Printf("Request completed! Status code %d\n", status)
	case <-ctx.Done():
		fmt.Printf("Request not completed Reason: %s!\n", ctx.Err())
	}

	fmt.Print("Process completed")
}

func doAsyncRequest(ch chan<- int, req *http.Request, client *http.Client) {
	res, err := client.Do(req)
	if err != nil {
		ch <- 0
		return
	}

	ch <- res.StatusCode
}
