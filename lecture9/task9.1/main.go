package main

import "fmt"

func main() {

	var availableFruits int
	var fruitName string

	fmt.Print("Введите название фрукта: ")
	fmt.Scanf("%s\n", &fruitName)

	availableFruits = fruitMarket(fruitName)

	if availableFruits != 0 {
		fmt.Printf("%v: в наличии %v", fruitName, availableFruits)
	}
}

func fruitMarket(fruit string) int {

	stock := map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    1,
		"груши":    0,
	}

	availableFruits := stock[fruit]
	if availableFruits == 0 {
		fmt.Printf("%v: нет в наличии", fruit)
	}

	return availableFruits
}
