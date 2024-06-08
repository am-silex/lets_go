package main

import (
	"encoding/xml"
	"fmt"
)

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	tenat    string
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {

	l := contracts{}
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов",
	}
	l.List = append(l.List, c)
	res, err := xml.Marshal(l)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", string(res))
}
