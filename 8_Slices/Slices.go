package main

import "fmt"

//slice
//dynamic arrays
//most used construct
//most useful methods are available

func main() {
	//uninitialised slice is nil
	// var num []int
	// fmt.Println(num == nil)
	// fmt.Println(len(num))

	var num = make([]int, 2, 5)

	// fmt.Println(num == nil)
	// fmt.Println(cap(num))

	num = append(num, 1)
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 5)
	fmt.Println(num)
	fmt.Println(cap(num))
	fmt.Println(len(num))

	// 	num := []int{1, 2, 3}
	// 	fmt.Println(cap(num))
	// 	fmt.Println(len(num))
	// 	fmt.Println(num)
	// }
	// var num = make([]int, 0, 5)
	// num = append(num, 2)
	// var num2 = make([]int, len(num))

	// //copy function
	// copy(num2, num)
	// fmt.Println(num, num2)

	//Slice operator
	// var num = []int{1, 2, 3}
	// fmt.Println(num[1:2])

}
