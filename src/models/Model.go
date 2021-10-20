package models

import "fmt"

type Model struct {

}

func (m *Model) Foo() {
	fmt.Println("Model Foo()")
}
