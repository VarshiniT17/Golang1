package main

import "fmt"

// to iterate over data structures
func main() {
	// 	fmt.Println("Range function")
	// 	data := "Hello world"
	// 	for index, value := range data {
	// 		fmt.Printf("Number at index %d is %c\n", index, value)
	// 	}

	// 	nums := []int{6, 7, 8}

	// 	// for i := 0; i < len(nums); i++ {
	// 	// 	fmt.Println(nums[i])
	// 	// }

	// 	// sum := 0

	// 	for i, num := range nums {
	// 		// sum = sum + num
	// 		// fmt.Println(sum)
	// 		fmt.Printf("index %d,value %d\n", i, num)
	// 	}

	m := map[string]string{
		"fname": "john",
		"lname": "doe",
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
