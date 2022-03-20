package main

import (
	"fmt"
	"sync"
)

/*
	Eğer birden fazla goroutine'in bitmesini bekliyorsan waitgroup kullanabilirsin.
*/
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(3)
//	go func() {
//		defer wg.Done()
//		fmt.Println("first")
//	}()
//	go func() {
//		defer wg.Done()
//		fmt.Println("second")
//	}()
//	go func() {
//		defer wg.Done()
//		fmt.Println("third")
//	}()
//	wg.Wait()
//}
//
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
		wg.Wait()
	}

	fmt.Println("test")
}
