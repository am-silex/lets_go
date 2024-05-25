package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	errorMessage string
}

func (me myFirstError) Error() string {
	return me.errorMessage
}

func main() {
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err1)
	err3 := fmt.Errorf("ошибка3:%w", err2)

	me := myFirstError{errorMessage: "Ошибка внтури структуры"}

	if errors.Is(err3, me) {
		fmt.Printf("Была ошибка %s", me)
	}
}
