package main

import (
	"fmt"
	// "time"
)

//Sending the data into the channels

// func processNum(numChan chan int) {

// 	for num := range numChan {

// 		fmt.Println("processing number", num)

// 		time.Sleep(time.Second * 1)
// 	}

// }
///////////to receive the data into the channels
// func sum(result chan int, num1 int, num2 int) {

// 	numresult := num1 + num2
// 	result <- numresult

// }

/////goroutine synchronizer
// func task(done chan bool) {

// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Processing.....")

// }

////////////////////////////queue system

// func emailSender(emailchan chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	////////////infinite loop ///////////
// 	for email := range emailchan {
// 		fmt.Println("sending email to", email)

// 		time.Sleep(time.Second * 1)

// 	}
// }

func main() {
	////////////for multiple channels processing we need select
	//////////select acts similar to switch
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {

		chan1 <- 10
	}()
	go func() {
		chan2 <- "pong"
	}()
	///////for more channels present
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received dat from chan1 :", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	///////////////////queue function sender
	// emailchan := make(chan string, 12)

	// done := make(chan bool)

	// go emailSender(emailchan, done)

	// for i := 0; i < 12; i++ {

	// 	emailchan <- fmt.Sprintf("%d@gmail.com ", i)

	// }

	// fmt.Println("done sending")
	// close(emailchan)

	// <-done
	// ///////////////////////////////////////////////////////////////
	// emailchan <- "1@bhemmappa.com"
	// emailchan <- "2ramappa@mail.com"

	// fmt.Println(<-emailchan)
	// fmt.Println(<-emailchan)

	//////////Synchronization
	// done := make(chan bool)/////unbuffered channel
	////////this channel has sending and receiving is blocking

	// go task(done)

	// <-done //// blocking////unbuffered and slow

	////use case of channels in go
	///for synchronisation
	////in place of the wait groups

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println("result", res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {

	// 	numChan <- rand.Intn(100)
	// }

	// numChan <-5
	// time.Sleep(time.Second * 2)

	// 	messageChannel := make(chan string)

	// 	messageChannel <- "ping"//channel are blocking until they are received

	// 	msg := <-messageChannel

	// 	fmt.Println(msg)

	var integerchannel chan int
	integerchannel = make(chan int)

}
