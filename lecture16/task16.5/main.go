package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				fmt.Println("сложные вычисления горутины: ", i)
				time.Sleep(time.Second)
				if ctx.Err() != nil {
					fmt.Println("stop горутина: ", i)
					return
				}
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		cancel()
	}()
	wg.Wait()
}
