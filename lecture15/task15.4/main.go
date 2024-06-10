package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func start() {
	fmt.Println("Running function start")
}

func main() {
	for j := 1; j <= 10; j++ {
		go func() {
			once.Do(start)
			fmt.Printf("Running goroutine %v\n", j)
		}()
	}
	time.Sleep(5 * time.Second)
}
