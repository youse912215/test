package main

import "fmt"

type Human struct {
	name  string
	level int
}

func (h *Human) initProcess(n string, l int) {
	h.name = n
	h.level = l
}

func (h *Human) updateHuman() (string, int) {
	return h.name, h.level
}

func main() {
	var h Humans
	name := "no name"
	level := 0

	fmt.Println(name, ":", level)

	h.initProcess("Mike", 9)
	fmt.Println(h.name, ":", h.level)
	name, level = h.updateHuman()
	fmt.Println("new human:", name, ":", level)
}
