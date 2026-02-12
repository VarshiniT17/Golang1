package main

import "fmt"

func main() {
	fmt.Println("Range function")
	data := "Hello world"
	for index, value := range data {
		fmt.Printf("Number at index %d is %c\n", index, value)
	}
}
