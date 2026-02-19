package main

import "fmt"

//enumerated types

type OrderStatus string

const (
	Received  OrderStatus = "received"
	confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "Delivered"
)

func changeOrderstatus(status OrderStatus) {
	fmt.Println("changing order status to ", status)
}

func main() {
	changeOrderstatus(Delivered)

}
