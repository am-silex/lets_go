package main

import "fmt"

func main() {
	a := 1

	do(a)
}
func do(v any) {
	a := 10

	v2, ok := v.(int)
	if !ok {
		fmt.Println("v is not int")
		return
	}
	a = a + v2

	fmt.Println(a)
}
