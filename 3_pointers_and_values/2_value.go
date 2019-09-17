package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	in10Seconds := now.Add(10 * time.Second)

	fmt.Println(now.Unix(), in10Seconds.Unix())
}

func main() {
	now := time.Now()
	in10Seconds := addTenSeconds(now)

	fmt.Println(now.Unix(), in10Seconds.Unix())
}

func addTenSeconds(t time.Time) time.Time {
	return t.Add(10 * time.Second)
}
