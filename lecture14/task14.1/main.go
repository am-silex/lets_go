package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()
	time.Sleep(3 * time.Second)
}
