package main

import "fmt"

func worker(i int) {
	fmt.Printf("worker %d Starting", 1)
	//asome task is happening
	fmt.Printf("worker %d Starting", 1)
}

func main() {

	//

	//Start 3 worker goroutines
	for i := 0; i <= 3; i++ {
		worker(i)
	}
	fmt.Println("workers task completedd")
}
