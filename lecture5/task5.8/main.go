package main

import "fmt"

type square int

func (s square) String() string {
	return fmt.Sprintf("%d м2", s)
}

func main() {

	var s square = 34
	s = s + square(10)

	fmt.Println(s)

}
