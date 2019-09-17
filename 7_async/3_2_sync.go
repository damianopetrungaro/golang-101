package main

import (
	"fmt"
	"sync"
)

type atomicCounter struct {
	i     int
	wg    *sync.WaitGroup
	mutex *sync.Mutex
}

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	c := &atomicCounter{0, wg, m}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go c.inc()
	}

	wg.Wait()
	fmt.Println(c.i)
}

func (c *atomicCounter) inc() {
	c.mutex.Lock()
	c.i++
	c.mutex.Unlock()
	c.wg.Done()
}
