package main

import "fmt"

func main() {

	type Origin1 string
	type Origin2 string
	var t1 Origin1 = "12:00"
	var t2 Origin2 = "17:59"

	fmt.Println("Afternoon:", t1, "Night", t2)

}
