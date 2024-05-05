package main

import "fmt"

func main() {

	var f func()

	f = func() {
		fmt.Print("Hello, Go!")
	}
	hello(f)
}

func hello(f func()) {
	f()
}
