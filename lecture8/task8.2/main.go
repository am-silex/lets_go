package main

import "fmt"

func main() {

	map1 := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	animals := [...]string{"слон", "бегемот", "осьминог"}

	for _, v := range animals {
		animalsCount, ok := map1[v]
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)", v, animalsCount, ok)
	}

}
