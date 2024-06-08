package main

import (
	"encoding/json"
	"fmt"
)

// В структуре поле tenat - не экспортируемое
// оно не будет выгружено в json
type tenant struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlor"`
	tenat    string `json:"tenat"`
}

func main() {
	t := tenant{Number: 1, Landlord: "Остап", tenat: "Шура"}

	fmt.Printf("%+v", t)
	fmt.Println()
	res, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(res))
	fmt.Println()
}
