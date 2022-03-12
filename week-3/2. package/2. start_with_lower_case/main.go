package main

import (
	"fmt"

	"github.com/h4yfans/go-patika.dev/formatter"
	"github.com/h4yfans/go-patika.dev/math"
)

func main() {
	num := math.double(2)
	output := formatter.format(num)
	fmt.Println(output)
}
