package main

//import "errors"

//
//import (
//	"errors"
//	"fmt"
//)

//func main() {
//	fmt.Println(sum(99, 1))
//}

//func sum(a, b int) int {
//	return a + b
//}

//---------//

//func addTo(base int, vals ...int) []int {
//	out := make([]int, 0, len(vals))
//	for _, v := range vals {
//		out = append(out, base+v)
//	}
//	return out
//}
//
//func main() {
//	fmt.Println(addTo(3))
//	fmt.Println(addTo(3, 2))
//	fmt.Println(addTo(3, 2, 4, 6, 8))
//	a := []int{4, 3}
//	fmt.Println(addTo(3, a...))
//	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
//}

//---------//
//func divide(numerator, denominator float64) (float64, error) {
//	if denominator == 0 {
//		return 0, errors.New("cannot divide by zero")
//	}
//
//	return numerator / denominator, nil
//}

//func main() {
//	result, err := divide(5, 3)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	resultDiv, _ := divide(5, 3)
//	fmt.Println(resultDiv)
//
//	fmt.Println(result)
//	result, err = divide(10, 2)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//
//	result, err = divide(8, 0)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//}
//
////---------//
//
//type divider func(int, int) int
//
//var opMap = map[string]divider{}

//---------//

//-----------//
//type Person struct {
//	FirstName string
//	LastName  string
//	Age       int
//}
//
//func main() {
//	people := []Person{
//		{"Zeytin", "Home", 21},
//		{"Ahmet", "Home", 21},
//		{"Fatma", "Home", 21},
//	}
//
//
//	sort.Slice(people, func(i int, j int) bool {
//		return people[i].LastName < people[j].LastName
//	})
//}
