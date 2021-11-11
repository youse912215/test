package main

import "fmt"

func main() {

	type Origin1 string
	type Origin2 string
	var t1 Origin1 = "12:00"
	var t2 Origin2 = "17:59"

	fmt.Println("Afternoon:", t1, "Night:", t2)

	array1 := [4]string{}
	array1[0] = "Spring"
	array1[1] = "Summer"
	array1[2] = "Automn"
	array1[3] = "Winter"
	fmt.Println(array1[0], array1[1], array1[2], array1[3])

	array2 := [...]string{"Apple", "Banana", "Grape"}
	fmt.Println(array2[0], array2[1], array2[2])

	array3 := []int{}
	for i := 0; i < 10; i++ {
		array3 = append(array3, i)
		fmt.Println(len(array3), cap(array3))
	}
}
