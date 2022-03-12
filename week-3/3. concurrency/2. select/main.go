package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
	select
		- Go'nun concurrency modelinin bir parçasıdır.
			Eğer birden fazla concurrent operasyonu yürütüyorsan ve önceliklendirmek istiyorsan bunun için select kullanmak gerekir.
*/

//import (
//	"context"
//	"fmt"
//	"log"
//	"time"
//)

//func writeFirstChan(ch chan string) {
//	time.Sleep(6 * time.Second)
//	ch <- "from channel 1"
//}
//func writeSecondChan(ch chan string) {
//	time.Sleep(3 * time.Second)
//	ch <- "from channel 2"
//
//}
//func main() {
//	output1 := make(chan string)
//	output2 := make(chan string)
//	go writeFirstChan(output1)
//	go writeSecondChan(output2)
//
//	for {
//		select {
//		case s1 := <-output1:
//			fmt.Println(s1)
//		case s2 := <-output2:
//			fmt.Println(s2)
//		}
//	}
//}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	check(ctx)

	fmt.Println("finale write")
}

func check(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			writeMessage()
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}
	}
}

func writeMessage() {
	log.Println("this is message")
}
