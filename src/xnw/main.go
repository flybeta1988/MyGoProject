package main

import (
	"../dbpkg"
	"fmt"
	"xnw.com/core"
)

func main() {
	core.Test()
	result := core.Join()
	fmt.Println("join result:", result)
	app := &core.App{}

	dbpkg.Foo()

	fmt.Println("app type:%T", app)
}
