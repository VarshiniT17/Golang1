package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}

}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
}

///ususally whenevr there are variables created in the functions they are of two types
//inner scope and outer scope
//the inner scope variables are created and then are removed from the memory once the program is executed
//but the outer scope variables keep track of the previously available result and also adds up the current result with same
