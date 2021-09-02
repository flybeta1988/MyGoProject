package core

import "fmt"

type App struct {
	id string
}

func (app *App) Foo() {
	fmt.Println("App's Foo func!")
}

func Test() {
	fmt.Println("App's Test func!")
}
