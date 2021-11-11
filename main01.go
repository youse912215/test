package main

import "fmt"

func main() {
	var (
		aaa int     = 123
		bbb float64 = 456
	)

	var name, age = "Yamada", 26

	const foo = 1000

	const (
		foo2 = 10000
		baa  = 39
	)

	fmt.Println("Hello, world")
	fmt.Println("aaa =", aaa, "bbb =", bbb)
	fmt.Println("name =", name, "age =", age)
}
