package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, deadline := context.WithDeadline(ctx, time.Now().Add(time.Second*2))
	defer deadline()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Goroutine done:", i, ctx.Err())
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
