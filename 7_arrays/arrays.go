package main

//numbered sequence of specific length
import "fmt"

func main() {
	// array length printing
	// 	var nums [5]int
	// 	fmt.Println(len(nums))

	// 	nums[0] = 9
	// 	fmt.Println(nums[0])

	// 	fmt.Println(nums)
	//

	// var vals [2]bool

	// fmt.Println(vals)

	var name [3]string

	name[0] = "r"
	name[1] = "a"
	name[2] = "m"
	fmt.Println(name)

	//zeroed values
	//int_> o
	//string_>""
	//bool-> false

	//to declare it in single line
	// 	num := [3]int{1, 2, 3}
	// 	fmt.Println(num)

	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums)

	//fixed size,that is know or predictable
	//memory optimization
	//constant time access
}
