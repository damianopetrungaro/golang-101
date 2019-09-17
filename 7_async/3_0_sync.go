package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	for _, v := range [5]string{"A", "B", "C", "D", "E"} {
		printAndDone(v, wg)
	}
	wg.Wait()
}

func printAndDone(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}
