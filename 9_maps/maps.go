package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map
	// m := make(map[string]string)

	// //setting an element
	// m["name"] = "go"
	// m["area"] = "backend"

	// //get an elemnet
	// fmt.Println(m["name"], m["area"])

	// //IMP: if key does not exists in the map then it returns zero value or null value
	// fmt.Println(m["phone"])

	//create map without using make keyword

	// fmt.Println(m)

	// m := map[string]int{"price": 47, "age": 30}

	// k, ok := m["price"]

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// fmt.Println(k)

	m1 := map[string]int{"price": 47, "age": 30}
	m2 := map[string]int{"price": 47, "age": 30}

	fmt.Println(maps.Equal(m1, m2))
}
