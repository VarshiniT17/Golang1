package main

import "fmt"

///copied by value
// func changenum(num int) {
// 	num = 5
// 	fmt.Println("in changenum", num)
// }

//by refrence
func changenum(num *int) {
	//de refer
	*num = 5
	fmt.Println("in changenum", *num)
}

func main() {
	///this variable has some value which will never change
	///the value doesnot change for the variables present in the main function even if you change it in other functions

	// even if the value changes for the same variable in other function
	//
	num := 1

	// changenum(num)
	changenum(&num)
	fmt.Println("after in chnangenum ", num)

	// fmt.Println("memory address", &num)
	// fmt.Println("after in chnangenum ", num)
	//to change the value of the variables present in the main function using the same variables present in the other functions we pass the reference of the variable
	//in the main function to the variable in the different function (address of the variable of the main function)

}
