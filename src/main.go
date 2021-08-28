package main

import (
	"fmt"
	"modtest/libs/dbs"
)

func main() {
	str := dbs.Test()
	fmt.Println(str)
	//dbpkg.Insert("flybeta")
}
