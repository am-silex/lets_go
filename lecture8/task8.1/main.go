package main

import "fmt"

func main() {

	map1 := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	fmt.Printf("%v", map1)

}
