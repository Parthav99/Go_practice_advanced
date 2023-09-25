package main

//-----------------------------------------------------------------------------------------------//
//Sending and receiving data from channel
//-----------------------------------------------------------------------------------------------//

// import "fmt"

// func main() {
// 	dataChannel := make(chan int, 1) //buffered channel

// 	// dataChannel := make(chan int) //unbuffered channel
// 	// go func() {
// 	// 	dataChannel <- 229
// 	// }()

// 	dataChannel <- 229
// 	n := <-dataChannel

// 	fmt.Println("Received from channel:", n)
// }

//-----------------------------------------------------------------------------------------------//
//Sending and receiving multiple data from channel
//-----------------------------------------------------------------------------------------------//

// import "fmt"

// func main() {

// 	// send is done with the help of background go routine
// 	// receive is done on the main go routine
// 	// we make use of an unbuffered channel
// 	// close after sending to channel, if ignored causes deadlock.
// 	// Because it keeps receiving, even when there is nothing to send.

// 	numberChannel := make(chan int)
// 	go func() {
// 		//sends to channel
// 		for i := 1; i <= 10; i++ {
// 			numberChannel <- i
// 		}
// 		close(numberChannel)
// 	}()

// 	//receives from channel
// 	for n := range numberChannel {
// 		fmt.Println(n)
// 	}
// }

//-----------------------------------------------------------------------------------------------//
//Sending random integers into channel, by creating multiple goroutines.
//-----------------------------------------------------------------------------------------------//

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"math/big"
// 	"sync"
// 	"time"
// )

// func doWork() *big.Int {
// 	time.Sleep(1 * time.Second)
// 	randomInts, _ := rand.Int(rand.Reader, big.NewInt(4))
// 	return randomInts
// }

// func main() {

// 	numChan := make(chan *big.Int)
// 	go func() {
// 		wg := sync.WaitGroup{} //to track running go routines
// 		for i := 0; i <= 5; i++ {
// 			wg.Add(1) //to increment everytime a go routine is created
// 			go func() {
// 				defer wg.Done() //decrement once the go routine id finished
// 				result := doWork()
// 				numChan <- result
// 			}()
// 		}
// 		wg.Wait() //waits until wait group counter is zero
// 		close(numChan)
// 	}()

// 	for n := range numChan {
// 		fmt.Println(n)
// 	}
// }

//-----------------------------------------------------------------------------------------------//
