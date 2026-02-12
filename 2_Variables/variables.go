package main

import "fmt"

func main() {
	// var name string = "golang"
	// var name = "golang" //when type is not mentioned , golanguage will  automatically infer the type based on the value assigned to it
	// var isAdult = true
	name := "golang"
	// fmt.Println(name)
	// fmt.Println(isAdult)

	//in some cases we might want to declare type of variable without assigning any value to it. In such cases we can use var keyword without assigning any value to it. The default value of the variable will be zero value of the type.
	// var name string
	fmt.Println(name) //output will be empty string as default value of string is empty string
}
