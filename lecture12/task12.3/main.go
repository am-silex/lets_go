package main

import "fmt"

type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данныевформате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данныевформате csv")
}

func main() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}

func Report(x interface{}) {
	switch v := x.(type) {
	case Xml:
		v.Format()
	case Csv:
		v.Format()
	}
}
