package main

import "fmt"

func main() {
	var name string
	fmt.Println("User mention your name :")

	fmt.Scan(&name)
	fmt.Println("Hello User ", name)
}
