package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var i int32
	i = 5
	go func() {
		atomic.AddInt32(&i, 1)
		fmt.Printf("atomic addition: %v\n", i)
	}()
	go func() {
		atomic.AddInt32(&i, -1)
		fmt.Printf("atomic subtraction: %v\n", i)
	}()
	time.Sleep(3 * time.Second)
}
