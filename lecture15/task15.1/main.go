package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		fmt.Println("Горутина: 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Горутина: 2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Горутина: 3")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Горутина: 4")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Горутина: 5")
	}()
	wg.Wait()
}
