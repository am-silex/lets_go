package main

import "fmt"

func main() {
	var numerator int32 = 16
	var denominator int32 = 3

	result := numerator / denominator
	remainder := numerator % denominator

	fmt.Printf("Результат: %v, остаток от деления: %v, тип результата: %T",
		result, remainder, result)
}
