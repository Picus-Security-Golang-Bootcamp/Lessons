package main

import (
	"fmt"
)

//----Goroutines, for Loops, and varying variables-----//
//https://github.com/golang/gofrontend/blob/e387439bfd24d5e142874b8e68e7039f74c744d7/go/statements.cc#L5501
func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	//for _, v := range a {
	//	go func() {
	//		ch <- v * 2
	//	}()
	//}

	//for _, v := range a {
	//	k := v
	//	go func() {
	//		ch <- k * 2
	//	}()
	//}

	for _, v := range a {
		go func(v int) {
			ch <- v * 2
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
