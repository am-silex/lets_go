package main

import "fmt"

type square int

func main() {

	var s square = 30
	s = s + square(15)

	fmt.Println(s)

}
