package main

import (
	"fmt"
)

type Human struct {
	age int
}

func (h *Human) walk() {
	fmt.Println("шлёп-шлёп")
}

type Action struct{
	Human Human
}

func (a *Action) run() {
	fmt.Println("бегу-бегу")
}

func main() {
	var Action Action
	Action.Human.age = 10
	Action.run()
	Action.Human.walk()
	return
}