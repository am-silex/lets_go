package main

// Разместил в папке pkg как рекомендуют
// Нужно ли создавать для каждого пакета отдельную подпапдку?

import (
	"fmt"
	first "lets_go/lecture10/task10.1/pkg"
)

func main() {
	fmt.Println(first.Hello())
}
