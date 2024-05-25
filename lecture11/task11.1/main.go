package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("the first error")
	err2 := fmt.Errorf("the second error: %w", err1)
	err3 := fmt.Errorf("and the third error: %w", err2)

	fmt.Println(err3)
}
