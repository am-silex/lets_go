package main

import "fmt"

// *Внутреннийсрезнаключе 1 долженостановитьвсециклы, начинаясовторогоцикла
// При такой постановке результат выполнения отличается от примера - в 4-м цикле так будет 1 строка, вместо 2
func main() {

	var s = []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		fmt.Printf("v1:%v\n", s[i])
		for i := 0; i < 1; i++ {
			fmt.Printf("\tv2:%v\n", s[i])
			for i := 0; i < 1; i++ {
				fmt.Printf("\t\tv3:%v\n", s[i])
				for i := 0; i < 1; i++ {
					fmt.Printf("\t\t\tv4:%v\n", s[i])
				}
			}
		}
	}
}
