package main

import (
	"fmt"
	"time"
)

//--Deadlock--- //
//Bir grup goroutine'in birbirini beklediğinde ve hiçbiri ilerleyemediğinde bir
//kilitlenme meydana gelir.
//Channellar gorotuinler arasında çalışmak içindir.
//func main() {
//	c1 := make(chan int)
//	fmt.Println("push c1: ")
//	c1 <- 10
//	g1 := <-c1
//	fmt.Println("get g1: ", g1)
//}

//func main() {
//	c1 := make(chan int)
//
//	go func() {
//		g1 := <-c1 // wait for value
//		fmt.Println("get g1: ", g1)
//	}()
//
//	fmt.Println("push c1 ")
//	c1 <- 10 // send value and wait until it is received.
//
//	time.Sleep(1 * time.Second)
//}

func main() {
	ch := make(chan int, 1)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
		// luck
		for i := 0; i < 100000; i++ {
			fmt.Println("a", i)
		}
	}()

	fmt.Println("here 1")
	fmt.Println("here 2")
	fmt.Println("here 3")
	fmt.Println(<-ch)
}
