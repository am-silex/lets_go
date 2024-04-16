package main

import "fmt"

const m = 5

func main() {

	var v *int
	f := func(v int) *int {
		return &v
	}
	v = f(m)
	v = f(4)

	fmt.Println(v, *v)

}
