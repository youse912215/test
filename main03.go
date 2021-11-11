package main

import "fmt"

func culc(x int, y int, v int) int {
	return x*y + v
}

func main() {

	var x int = 5
	max := 2

	fmt.Println("Before:", x)

	if x > max {
		x = max
		fmt.Println("After:", x)
	}

	y := 8

	for x < y {
		x++
		fmt.Println("for:", x)
		if x == 6 {
			fmt.Println("break")
			break
		}
	}

	add1 := 2
	add2 := 3

	x = culc(add1, add2, x)
	fmt.Println("Add:", x)
}
