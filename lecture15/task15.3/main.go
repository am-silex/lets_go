package main

import (
	"fmt"
	"sync"
	"time"
)

func MWriter(c *int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	*c++
	time.Sleep(time.Second)
}

func MReader(c *int, mu *sync.Mutex) {
	mu.Lock()
	fmt.Println(*c)
	mu.Unlock()
	time.Sleep(time.Second)
}

func main() {
	var c int
	mu := new(sync.Mutex)
	go MReader(&c, mu)
	for j := 0; j < 100; j++ {
		go MWriter(&c, mu)
		go MReader(&c, mu)
	}
	time.Sleep(10 * time.Second)
}
