package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

//
//func getSiteContent(id int, url <-chan string, result chan<- string) {
//	for v := range url {
//		resp, _ := http.Get(v)
//		//defer resp.Body.Close()
//		resp.Body.Close()
//		respBodyBytes, _ := ioutil.ReadAll(resp.Body)
//		fmt.Printf("ID : %d, URL : %s\n", id, v)
//		result <- string(respBodyBytes)
//
//		fmt.Printf("Sent to channel : %s \n", v)
//	}
//}
//
//func main() {
//	urls := []string{
//		"https://yolcu360.com",
//		"https://adjust.com",
//		"https://google.com",
//		"https://twitter.com",
//		"https://youtube.com",
//		"https://yahoo.com",
//	}
//
//	url := make(chan string, len(urls))
//	result := make(chan string, len(urls))
//
//	for i := 0; i < 3; i++ {
//		go getSiteContent(i, url, result)
//	}
//
//	for i := 0; i < len(urls); i++ {
//		//time.Sleep(2 * time.Second)
//		url <- urls[i]
//	}
//
//	close(url)
//
//	for i := 0; i < len(urls); i++ {
//		a := <-result
//		fmt.Println(a)
//	}
//}

func getSiteContent(wg *sync.WaitGroup, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		wg.Done()
		resp.Body.Close()
	}()
	_, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("URL : %s\n", url)
	fmt.Println("data")
}

func main() {
	urls := []string{
		"https://yolcu360.com",
		"https://adjust.com",
		"https://google.com",
		"https://twitter.com",
		"https://youtube.com",
		"https://yahoo.com",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			getSiteContent(&wg, url)
		}(url)
	}
	wg.Wait()

	fmt.Println("son")
}
