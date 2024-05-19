package main

import "fmt"

func main() {

	var foodName string
	fmt.Printf("Введите название еды: ")
	fmt.Scanf("%s\n", &foodName)
	checkFood(foodName)
}

func checkFood(checkName string) {

	fruitOrVegetable := map[string]bool{
		"груша":    true,
		"яблоко":   true,
		"апельсин": true,
		"тыква":    false,
		"огурец":   false,
		"помидор":  false,
	}

	switch v, ok := fruitOrVegetable[checkName]; {
	case ok && v == true:
		{
			fmt.Printf("Это фрукт")
		}
	case ok && v == false:
		{
			fmt.Printf("Это овощь")
		}
	default:
		{
			fmt.Printf("Что-то странное...")
		}

	}
}
