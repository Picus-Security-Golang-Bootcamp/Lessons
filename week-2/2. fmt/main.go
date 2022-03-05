package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	args := os.Args[1:]

	for i, arg := range args {
		fmt.Printf("index: %d, value: %s\n", i, arg)
	}
}
