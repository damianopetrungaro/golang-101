package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	ctx, _ = context.WithTimeout(ctx, 650*time.Millisecond)
	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response retrieved!", res.StatusCode)
}
