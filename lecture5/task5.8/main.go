package main

import "fmt"

type square int

func main() {

	var s square = 34
	s = s + square(10)

	fmt.Println(s, "m2")
	// Не нашел, как реализовать функцию сериализации для пользовательского типа
	// для автоматиеческого форматирования при любом выводе значения

}
