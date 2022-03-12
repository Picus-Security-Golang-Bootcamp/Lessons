package main

import "fmt"

//----Goroutine leaks -----/
//https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html
//https://groups.google.com/g/golang-nuts/c/pZwdYRGxCIk/m/qpbHxRRPJdUJ
// It's only necessary to close a channel when it is important to tell the
//receiving goroutines that all data have been sent
func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	//for i := range countTo(10) {
	//	fmt.Println(i)
	//}

	for i := range countTo(10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
