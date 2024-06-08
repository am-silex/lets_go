package main

import (
	"encoding/json"
	"fmt"
)

// Структура имеет 2 поля со строчной буквы - эти  поля будут приватным
// и не будут заполнятся прочитанными данными из строки json
type tenant struct {
	Number   int `json:"number"`
	Landlord string
	Tenat    string `json:"tenat"`
}

func main() {
	InputData := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	t := tenant{}
	err := json.Unmarshal([]byte(InputData), &t)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", t)
	fmt.Println()
}
