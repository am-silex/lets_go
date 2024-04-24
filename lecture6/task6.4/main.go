package main

import "fmt"

type user struct {
	ID   int
	Name string
	contacts
}
type employee struct {
	ID   int
	Name string
	contacts
}

type contacts struct {
	Addss string
	Phone string
}

func main() {

	u := user{
		ID:   1,
		Name: `Alexey`,
		contacts: contacts{
			Addss: "Somewhere on the planet",
			Phone: "555",
		},
	}

	e := employee{
		ID:   1,
		Name: "Employee Alexey",
		contacts: contacts{
			Addss: "In the office",
			Phone: "495",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)

}
