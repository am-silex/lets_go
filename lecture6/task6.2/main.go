package main

import "fmt"

func main() {

	type contract struct {
		ID     int
		Number string
		Date   string
	}

	c := contract{
		ID:     1,
		Number: "##000A101\t01",
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v", c)

}
