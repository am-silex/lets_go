package main

// Разместил в папке pkg как рекомендуют
// Нужно ли создавать для каждого пакета отдельную директорию?

import (
	"fmt"
	first "github.com/am-silex/lets_go/lecture10/task10.1/pkg"
)

func main() {
	fmt.Println(first.Hello())
}
