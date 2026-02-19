package main

import "fmt"

// func printslice[T interface{}](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printslice[T int | string | bool | float32](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printstringslice[T any](items []T) {
// 	for _, item := range items {

// 		fmt.Println(item)
// 	}
// }

//LIFO
// type stack[T any] struct {
// 	elements []T
// }

func printslice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	// myStack := stack[int]{
	// 	elements: []int{1, 2, 3},
	// }

	// nums := []int{1, 2, 3}
	vals := []bool{true, false, true}

	// names := []string{"golang", "varshini", "varun"}
	// printslice(names)
	// printslice(nums)
	printslice(vals)

	// fmt.Println(myStack)
}
