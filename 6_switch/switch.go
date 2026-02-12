package main

import "fmt"

func main() {
	//simple switch
	// i := 2

	// switch i {
	// case 1:
	// 	fmt.Println("go")
	// case 2:
	// 	fmt.Println("goo")
	// default:
	// 	fmt.Println("other")
	// }

	//switch with multiple conditions

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")
	// default:
	// 	fmt.Println("its weekday")
	// }

	//type switch
	WhoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other ", t)
		}
	}

	WhoAmI(33.45)
}
