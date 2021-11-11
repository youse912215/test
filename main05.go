package main

import "fmt"

/* Polymorphism */

type Human struct {
	name  string
	level int
}

type Item struct {
	name string
}

func (h Human) showName() string {
	return h.name
}

func (h Human) copyFunc() {
	fmt.Println(h.showName())
}

func (i Item) showName() string {
	return i.name
}

func (i Item) copyFunc() {
	fmt.Println(i.showName())
}

func main() {
	h := Human{name: "Kevin"}
	i := Item{name: "Rocket Punch"}
	h.copyFunc()
	i.copyFunc()
}
