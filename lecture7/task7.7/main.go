package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	var resultSlice []int
	resultSlice = append(resultSlice, slice1...)
	resultSlice = append(resultSlice, slice2...)

	fmt.Printf("%v", resultSlice)

}
