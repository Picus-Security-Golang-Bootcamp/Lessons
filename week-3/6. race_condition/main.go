package main

import (
	"fmt"
	"sync"
)

// Birden fazla goroutine'in aynı memory alanına aynı anda erişilmesi istenmedi
//noktada mutex kullanılır.
// Concurrent read/write problemlerini çözerken mutex kullanılırz.
//
//func main() {
//	m := map[int]int{}
//	go writeLoop(m)
//	go readLoop(m)
//
//	// stop program from exiting, must be killed
//	block := make(chan struct{})
//	<-block
//}
//
//// writeLoop continuosly mutates the value
//func writeLoop(m map[int]int) {
//	for {
//		for i := 0; i < 100; i++ {
//			m[i] = i
//		}
//	}
//}
//
//// readLoop prints the value
//func readLoop(m map[int]int) {
//	for {
//		for k, v := range m {
//			fmt.Println(k, "-", v)
//		}
//	}
//}

//----Mutex---- //
//func main() {
//	m := map[int]int{}
//
//	mux := &sync.Mutex{}
//
//	go writeLoop(m, mux)
//	go readLoop(m, mux)
//
//	// stop program from exiting, must be killed
//	block := make(chan struct{})
//	<-block
//}
//
//func writeLoop(m map[int]int, mux *sync.Mutex) {
//	for {
//		for i := 0; i < 100; i++ {
//			mux.Lock()
//			m[i] = i
//			mux.Unlock()
//		}
//	}
//}
//
//func readLoop(m map[int]int, mux *sync.Mutex) {
//	for {
//		mux.Lock()
//		for k, v := range m {
//			fmt.Println(k, "-", v)
//		}
//		mux.Unlock()
//	}
//}

//--------RWMutex----------//
func main() {
	m := map[int]int{}

	mux := &sync.RWMutex{}

	go writeLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock()
			m[i] = i
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.RUnlock()
	}
}
