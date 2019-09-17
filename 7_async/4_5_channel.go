package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	defer close(done)

	go slowWork(done)
	jobDone(done)
}

func jobDone(done <-chan bool) {
	<-done
	fmt.Println("Done!")

}

func slowWork(done chan<- bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("After 2 seconds I did my job!")
	done <- true
}
