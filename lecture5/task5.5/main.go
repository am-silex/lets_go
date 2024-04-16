package main

import "fmt"

func main() {

	var s string = "Local variable"
	fmt.Println(s)
	changeLocalVariable(&s)
	fmt.Println(s)

}

func changeLocalVariable(vAdd *string) {
	*vAdd = "Changed local variable"
}
