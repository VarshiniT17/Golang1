// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func task(id int, w *sync.WaitGroup) {
// 	defer w.Done()
// 	fmt.Println("Doing task", id)
// }

// func main() {

// 	var wg sync.WaitGroup
// 	for i := 0; i <= 10; i++ {
// 		wg.Add(1)
// 		go task(i, &wg)

// 		// go func(i int) {
// 		// 	fmt.Println(i)
// 		// }(i)
// 	}

// 	wg.Wait()

// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	time.Sleep(1 * time.Second)
// }

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello world")
	// time.Sleep(500 * time.Millisecond)
	fmt.Println("hiiiii")
}

func sayHii() {
	fmt.Println("hello")
	//		// time.Sleep(1000 * time.Millisecond)
}
func main() {

	fmt.Println("learning go routine")

	go sayHello()

	go sayHii()
	time.Sleep(500 * time.Millisecond)
}
