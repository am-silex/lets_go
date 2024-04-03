package main

import "fmt"

func main() {

	var f func() string

	f = func() string {
		return fmt.Sprint("Hello, Go!")
	}
	hello(f())
}

func hello(f string) {
	fmt.Print(f)
}
