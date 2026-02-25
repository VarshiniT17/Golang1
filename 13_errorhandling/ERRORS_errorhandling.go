package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return a / b, nil

}
func main() {
	// fmt.Println("Error handling")

	data, _ := divide(10, 2)
	fmt.Println("answer= ", data)
}

// ////////////////////////////////////////////////////////////////////////
// // package main

// // import "fmt"

// // func sum(a, b int) (result int, err error) {

// // 	if a < 0 || b < 0 {
// // 		return 0, errors.New("values cannot be zero or less than zero")
// // 	}
// // 	return a + b, nil
// // }
// // func main() {
// // 	c, err := sum(10, -1)

// // 	if err != nil {
// // 		fmt.Println(err.Error())
// // 	}
// // }

// //In Go, errors are values.

// //Go does NOT use try-catch like Java/Python.
// //Instead, functions return an error as the last return value.
// // value, err := someFunction()
// // if err != nil {
// //     // handle error
// // }

// // The error Interface
// // In Go, error is just an interface:

// type error interface {
//     Error() string
// }
// Any type that implements Error() string is an error.

// Creating Errors
// ✅ Using errors.New()
// import "errors"

// err := errors.New("something went wrong")
// ✅ Using fmt.Errorf()
// import "fmt"

// err := fmt.Errorf("user %d not found", id)
