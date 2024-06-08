package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	return d.voice
}

func main() {
	var d = Duck{}
	d.voice = "Голос птички - решение №3"
	song, err := Sing(&d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	switch b.(type) {
	case nil:
		return "", errors.New("Ошибка пения!")
	default:
		return b.Sing(), nil
	}

}
