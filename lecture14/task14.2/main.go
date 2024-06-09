package main

import "fmt"

func main() {

	ch := make(chan string)
	go func() {
		message := "Привет, строковый канал!"
		ch <- message
	}()
	fmt.Println(<-ch)
}
