package main

import (
	"fmt"
	"math"
)

var (
	globalIntVar = 100
)

func main() {
	sum := globalIntVar + 50

	fmt.Printf("Total value : %d\n", sum)

	sumAbs := math.Abs(float64(sum))

	fmt.Printf("Sum Abs result : %v\n", sumAbs)

	fmt.Printf("%.3f", 12.3456)
}

// sum
func sum(a, b int) int {
	return 0
}
