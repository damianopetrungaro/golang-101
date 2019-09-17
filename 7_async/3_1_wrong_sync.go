package main

import (
	"fmt"
	"sync"
)

type counter struct {
	i  int
	wg *sync.WaitGroup
}

func main() {
	wg := &sync.WaitGroup{}
	c := &counter{0, wg}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go c.inc()
	}

	wg.Wait()
	fmt.Println(c.i)
}

func (c *counter) inc() {
	c.wg.Done()
	c.i++
}
