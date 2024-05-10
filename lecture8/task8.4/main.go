package main

import "fmt"

func main() {

	map1 := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	map1["выдра"] = struct{}{}

	fmt.Print(map1)

}
