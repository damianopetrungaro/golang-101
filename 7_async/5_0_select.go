package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // used to randomize next Run

	ch := make(chan string)
	go notify(ch)
	timeout := time.After(time.Second)

	for {
		select {
		case <-timeout:
			fmt.Println("one seconds passed!, Exiting")
			return
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}

func notify(ch chan<- string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- fmt.Sprintln("Message number", i)
	}
}
