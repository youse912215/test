package main

import "fmt"

int main() {

	var x int = 5
	const max = 2

	fmt.Println("Before:", x)

	if x > max {
		max = x
		fmt.Println("After:", x)
	}	
}