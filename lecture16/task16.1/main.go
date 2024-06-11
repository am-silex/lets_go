package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Goroutine %v done\n", j)
					return
				default:
					fmt.Println("Processing...")
				}
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Main goroutine stopped")
		cancel()
	}()
	wg.Wait()
}
