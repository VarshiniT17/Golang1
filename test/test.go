package main

import "fmt"

func add(new []int) []int {
	new = append(new, 5)
	new = append(new, 8)
	return new
}

func main() {
	new := make([]int, 2, 5)
	new[0] = 1
	new[1] = 3

	new = add(new)

	fmt.Println(new[0:3])
}
