package main

import (
	formater "github.com/am-silex/lets_go_module_format/v2"
)

func main() {

	fileSource := "lecture13//task13.5//db.txt"
	fileResult := "db_out.txt"

	formater.Do(fileSource, fileResult)

}
