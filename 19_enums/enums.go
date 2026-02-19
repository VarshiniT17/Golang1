package main

import "fmt"

// Define a custom type for order status
type OrderStatus string

// Enum-like constants for OrderStatus
const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

// Function to change order status
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

func main() {
	// Example usage
	changeOrderStatus(Received)
	changeOrderStatus(Prepared)
	changeOrderStatus(Delivered)
}
