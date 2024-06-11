package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Goroutine stopped: ", i, ctx.Err())
					return
				default:
					fmt.Println("Processing...")
					time.Sleep(time.Second)
				}
			}
		}()
	}
	wg.Wait()
}
