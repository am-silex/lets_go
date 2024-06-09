package main

import "fmt"

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	close(ch)
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	<-stop
	fmt.Println("happy end")
}
