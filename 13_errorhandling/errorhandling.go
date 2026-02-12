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
