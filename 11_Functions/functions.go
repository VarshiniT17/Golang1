package main

import "fmt"

// func simplefunction() {
// 	fmt.Println("simplefunction")
// }

func add(a, b int) (result int) {
	result = a + b
	return

}
func main() {

	// fmt.Println("learning functions  ")
	// simplefunction()

	result := add(10, 3)
	fmt.Println("sum =", result)
}
