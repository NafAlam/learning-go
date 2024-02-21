// // THIS HASNT BEEN FINISHED!!!....
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func add(x int, y int, returnChan chan int) {
// 	time.Sleep(1 * time.Second)
// 	returnChan <- x + y
// }

// func main() {
// 	// when you want to return a value from a go routine or indicate youve finished completing something use Channel
// 	returnChan := make(chan int)
// 	go add(1, 2, returnChan)
// 	go add(1, 5, returnChan)
// 	go add(1, 8, returnChan)
// 	go add(1, 9, returnChan)
// 	returnValue := <-returnChan // blocking code meaning the channel has to receive something
// 	fmt.Println(returnValue)
// 	returnValue = <-returnChan // blocking code meaning the channel has to receive something
// 	fmt.Println(returnValue)
// 	returnValue = <-returnChan // blocking code meaning the channel has to receive something
// 	fmt.Println(returnValue)
// 	returnValue = <-returnChan // blocking code meaning the channel has to receive something
// 	fmt.Println(returnValue)

// 	// these will not return in the same order as you have written the code

// 	chan1 := make(chan int)
// 	chan2 := make(chan int)

// 	go add(99, 5, chan1)
// 	go add(-10, -5, chan2)

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case rv1 := <-chan1:
// 			fmt.Println(rv1)
// 		case rv2 := <-chan2:
// 			fmt.Println(rv2)

// 		}
// 	}

// 	// this a buffer channel meaning it can hold 3 items
// 	ch := make(chan int, 3)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch) // a channel must receive a elemnet before it can display and be spat out
// 	<-ch              // this is blocking as it is waiting for somehting to send to the channel
// }
