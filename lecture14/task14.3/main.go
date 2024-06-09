package main

import "fmt"

func main() {

	ch := make(chan string, 4)
	go func() {
		ch <- "Привет, буферизованный канал!"
		ch <- "Привет, буферизованный канал!"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
