package main

import "fmt"

func main() {

	map1 := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	map1["бегемот"] = 2

	fmt.Print(map1)

}
