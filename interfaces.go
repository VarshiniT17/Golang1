// package main

// import "fmt"

// type paymenter interface {
// 	pay(amount float32)
// }

// type payment struct {
// 	gateway paymenter
// }

// //open close principle violation
// //open for extention but closed for modification

// func (p payment) makePayment(amount float32) {

// 	// razorpayPaymentGw := razorpay{}
// 	// razorpayPaymentGw.pay(amount)
// 	//
// 	p.gateway.pay(amount)

// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	//logic to make payment

// 	fmt.Println("making payment using razorpay", amount)
// }

// type paypal struct{}

// func (m paypal) pay(amount float32) {
// 	fmt.Println("making payment using paypal")
// }

// // type stripe struct{}

// // func (s stripe) pay(amount float32) {
// // 	fmt.Println("making payment using stripe", amount)
// // }

// type fakePayment struct{}

// func (f fakePayment) pay(amount float32) {
// 	fmt.Println("making payment using fake gateway for testing purpose")
// }

// func main() {

// 	// stripePaymentGw := stripe{}
// 	// razorpayPaymentGw := razorpay{}
// 	paypalPaymentGw := paypal{}

// 	// fakeGw := fakePayment{}
// 	newPayment := payment{
// 		gateway: paypalPaymentGw,
// 	}
// 	newPayment.makePayment(100)
// }
