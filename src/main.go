package main

import (
	"dbpkg"
	"fmt"
	"libs"
)

func main() {
	str := libs.FunA()
	fmt.Println(str)
	dbpkg.Insert("flybeta")
}
