package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("Animal speaking")
}

type Dog struct {
	Animal
	Breed string
}

// type home struct {
// 	name    string
// 	address string
// 	houseno int
// }

// func main() {

// 	var house home

// 	type home struct {
// 		name    string
// 		address string
// 		houseno int
// 	}

// 	fmt.Println("house:", house)
// 	house.name = "ram"
// 	house.address = "sharma"
// 	house.houseno = 312
// 	fmt.Println("house :", house)

// 	home1 := home{
// 		name:    "radha",
// 		address: "sham",
// 		houseno: 32,
// 	}

// 	ice := home{
// 		name:    "cone",
// 		address: "mango",
// 		houseno: 2,
// 	}
// 	fmt.Println("ice creame:", ice)
// 	fmt.Println("home1:", home1)
// }

// package main

func main() {

	d := Dog{
		Animal: Animal{Name: "Tommy"},
		Breed:  "Labrador",
	}

	d.Speak() // inherited method

	// type home struct {
	// 	name string
	// 	no   int
	// }

	// house := home{
	// 	name: "radha",
	// 	no:   32,
	// }
	// fmt.Println("house1:", house)

	// var home2 = new(home)
	// home2.name = "quintin tarantino"
	// home2.no = 21

	// fmt.Println("new entry ", home2)
}
