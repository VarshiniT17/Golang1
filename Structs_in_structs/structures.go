///nesteed structs or compossition structs

// package main

// import "fmt"

// type person struct {
// 	name    string
// 	address string
// }

// type contact struct {
// 	no    int64
// 	email string
// }

// type place struct {
// 	city  string
// 	state string
// }

// type employee struct {
// 	Person_details person
// 	Person_contact contact
// 	Person_place   place
// }

// func main() {

// 	employee1 := employee{
// 		Person_details: person{
// 			name:    "ram",
// 			address: "sharma",
// 		},
// 		Person_contact: contact{
// 			no:    8088135283,
// 			email: "varshiniturmari@gmail.com",
// 		},
// 		Person_place: place{
// 			city:  "dwd",
// 			state: "kar",
// 		},
// 	}

// 	fmt.Println("details:", employee1)
// }

// / embedding structs
package main

import "fmt"

type person struct {
	name    string
	address string
}

type contact struct {
	no    string
	email string
}

type place struct {
	city  string
	state string
}

type employee struct {
	person
	contact
	place
}

func main() {

	employee1 := employee{
		person: person{
			name:    "ram",
			address: "sharma",
		},
		contact: contact{
			no:    "8088135283",
			email: "varshiniturmari@gmail.com",
		},
		place: place{
			city:  "dwd",
			state: "kar",
		},
	}
	fmt.Println("employee:", employee1)
	// fmt.Println("Name:", employee1.name)
	// fmt.Println("City:", employee1.city)
	// fmt.Println("Email:", employee1.email)
}
