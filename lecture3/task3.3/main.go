package main

import "fmt"

const myPackageConst = 1

func main() {
	const myPackageConst = 2
	fmt.Print(myPackageConst)
}
