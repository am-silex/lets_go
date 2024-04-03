package main

import "fmt"

func main() {

	var v func()
	v = hello()
	v()

}

func hello() func() {

	return func() {
		fmt.Print("Hello, GO!")
	}

}
