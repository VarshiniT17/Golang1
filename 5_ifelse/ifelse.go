package main

import "fmt"

func main() {
	//age := 29

	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }

	//we can declare a variable inside a construct
	if age := 3; age >= 18 {
		fmt.Println("person is adult", age)
	} else if age >= 12 {
		fmt.Println("person is a teenager", age)
	} else {
		fmt.Println("person is a kid ", age)
	}

	// var role ="admin"
	// var hasPermissions =true

	// if role ="admin" && hasPermissions
	// {
	// 	fmt.Println("yes ")
	// }

	// 	bakery := "building"
	// 	Hasbuns := false

	// 	if bakery == "building" && Hasbuns {
	// 		fmt.Println("Customer can go to the bakery ")
	// 	} else {
	// 		fmt.Println("Customer can go to the bakery but cannot buy buns")
	// 	}
	//

	// go does not have ternary operator , you will have to use normal if else only
}
