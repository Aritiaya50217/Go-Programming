package main

import (
	"fmt"
)

// interface
type Phone interface {
	Call(number string)
}

type Samsung struct {
	Name string
}

func (s Samsung) Call(number string) {
	fmt.Println(s.Name, "Calling", number)
}

func (s Samsung) Answer() {
	fmt.Println(s.Name, "hello")
}

func Dial(p Phone) {
	p.Call("0987654321")
}

func main() {
	s := Samsung{
		Name: "S10",
	}
	s.Answer()

	Dial(s)
}
