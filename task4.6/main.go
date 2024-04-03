package main

import "fmt"

func main() {

	printSequenceNumber(0, 1, 0)

}

func printSequenceNumber(a, b, i int) {

	if i >= 23 {
		return
	}
	c := a + b
	i += 1
	fmt.Printf("%v - порядок, значение = %v \n", i, c)
	printSequenceNumber(b, c, i)

}
