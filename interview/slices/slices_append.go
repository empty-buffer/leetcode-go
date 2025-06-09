package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	println(slice)
	fmt.Println(slice)

	add(slice[0:2], 5)
	println(slice)
	fmt.Println(slice)

	println(slice)
	fmt.Println(slice)
	add(slice[0:2], 5, 6, 7)

	mm := map[string]int{"a": 1, "b": 4}
	for k, v := range mm {
		fmt.Println(k, v)
	}
}

func add(slice []int, num ...int) {
	for _, v := range num {
		slice = append(slice, v)
	}
	println(slice)
}
