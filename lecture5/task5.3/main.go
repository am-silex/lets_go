package main

import "fmt"

func main() {

	s := "Test string"

	var sPtr *string

	sPtr = &s

	fmt.Println("Variable before: ", s)
	*sPtr = "Changed string"
	fmt.Println("Variable after: ", s)

}
