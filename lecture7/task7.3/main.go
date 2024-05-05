package main

import "fmt"

func main() {

	m := [...]string{"яблоко", "груша", "помидор", "абрикос"}
	m[2] = "персик"
	fmt.Printf("%v", m)

}
