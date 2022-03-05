package main

func main() {
	// Boolean
	//var flag bool
	//fmt.Println(flag)

	//var patika bool = true
	//fmt.Println(patika)
	//patika = "patika"

	// Integer Types
	//var num int64 = 10

	// Float
	//var num1 float32 = 10.2
	//var num2 float64 = 10.2

	// var vs :=
	//var x int = 10
	//var x = 10
	//var x int
	//fmt.Println(x)

	//var x = 10
	//var y float64 = 20
	//y:=20

	//var x, y int = 10, 20
	//var x, y = 10, "patika"

	//var (
	//	x    int
	//	y        = 20
	//	z    int = 30
	//	d, e     = 40, "hello"
	//	f, g string
	//)

	//const
	//const pi = 3
	//fmt.Println(pi)
	//pi = 4
	//fmt.Println(pi)

	//const x float64 = 10
	//
	//const max_value = 1
	//const MAX_VALUE = 1
	//
	//const MAXVALUE = 1
	//const maxValue = 1

	//Composite Types
	//Arrays
	//var x [3]int
	//var y = [3]int{10, 20, 30}
	//y = append(y, 40)
	//var x = []int{10, 20, 30}
	//var x [][]int
	// Slices
	// [...] -> array
	// [] -> slice

	//x := []int{10, 20, 30}
	//fmt.Println(len(x))

	//var x []int
	//y := []int{40, 50, 60}
	//x = append(x, y...)
	//fmt.Println(x)

	//Capacity
	//var x []int
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 10)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 20)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 30)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 40)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 50)
	//fmt.Println(x, len(x), cap(x))

	//Make
	//x := make([]int, 5)
	//fmt.Println(x)
	//x = append(x, 99)
	//fmt.Println(x)

	//x := make([]int, 5, 10)
	//fmt.Println(x)

	// Slicing Slices
	//x := []int{1, 2, 3, 4}
	//y := x[:2]
	//fmt.Println(y)

	//z := x[1:]
	//fmt.Println(z)

	//e := x[:]
	//fmt.Println(e)

	//x := [4]int{1, 2, 3, 4}
	//y := x[:2]

	// copy
	//x := []int{1, 2, 3, 4}
	//y := x
	//y := make([]int, len(x))
	//num := copy(y, x)
	//fmt.Println(y, num)

	//Maps
	//var x map[string]int
	//totalWins := map[string]int{}
	//teams := map[string][]string{
	//	"Ahmet": []string{"a", ""},
	//	"Ayşe":  []string{"e", "f", "g"},
	//}
	//teams["Sevde"] = []string{"a", "y"}
	//teams["Sevde"] = []string{"x", "o"}
	////fmt.Println(teams)
	//
	//y := teams["Ayşe"]
	//y = append(y, "k")
	//teams["Ayşe"] = y
	//fmt.Println(teams)

	//ages := make(map[string]string, 10)
	//fmt.Println(ages)

	// Reading and Writing a Map

	//totalWins := map[string]int{}
	//totalWins["Ayşe"] = 1
	//totalWins["Osman"] = 2
	//totalWins["Ayşe"]++
	//fmt.Println(totalWins)

	// The comma ok idiom
	//m := map[string]string{
	//	"hello": "5",
	//	"world": "0",
	//}
	//v, ok := m["goodbye"]

	//fmt.Println(v, ok)
	//delete(m,"hello")
	//fmt.Println(m)

	// Struct
	//type person struct {
	//	name string
	//	age  int
	//	pet  string
	//}

	//var ahmet person
	//osman := person{
	//	"Osman",
	//	21,
	//}
	//ayse := person{}

	//
	//
	//osman := person{
	//	name: "Osman",
	//	age:  21,
	//	pet:  "Zeytin",
	//}
	//
	//osman.age = 22
	//osman.pet = "turuncu"
	//

	// anonymous structs

	//Struct
	//type person struct {
	//	name string
	//	age  int
	//	pet  string
	//}
	//
	//var person struct {
	//	name string
	//	age  int
	//	pet  string
	//}

	//person.name = "Osman"
	//person.age = 21
	//person.name = "Zeytin"
	//
	//pet := struct {
	//	name string
	//	kind string
	//}{
	//	name: "Zeytin",
	//	kind: "Cat",
	//}
	//

	//
	//type firstPerson struct{
	//	name string
	//	age int
	//}
	//f = firstPerson{
	//	name: "Osman",
	//	age:  21,
	//}
	//

	//type person struct {
	//	name string
	//	age  int
	//	pet  string
	//}
	//var g struct {
	//	name string
	//	age  int
	//}
	//g.age = 21
	//g.name = "test"
	//g = f
	//
	//
	//fmt.Println(f == g)
	//

	// Shadowing Variables

	//x := 10
	//if x > 5 {
	//	fmt.Println(x)
	//	x = 5
	//	fmt.Println(x)
	//}
	//fmt.Println(x)

	// if
	//n := 10
	//if n == 0{
	//	//
	//}else if n == 1{
	//
	//}else{
	//
	//}

	// for
	// c-style
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	////condition-only
	//i := 1
	//for i < 100 {
	//	fmt.Println(i)
	//	i = i * 2
	//}
	//
	//// infinite
	//for {
	//	fmt.Println("HELLO")
	//}
	//
	//break
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	if i == 5{
	//		fmt.Println("finished")
	//		break
	//	}
	//}

	//continue
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	if i == 5{
	//		fmt.Println("finished")
	//		continue
	//	}
	//}

	//for-range
	//evenVals := []int{2, 4, 6, 8, 10}
	//for i, v := range evenVals {
	//	fmt.Println(i,v)
	//}

	//evenVals := []int{2, 4, 6, 8, 10}
	//for _, v := range evenVals {
	//	fmt.Println(v)
	//}

	// iterate
	//m := map[string]int{
	//	"a": 1,
	//	"c": 3,
	//	"b": 2,
	//}
	//
	//for i := 0; i < 3; i++ {
	//	fmt.Println(m)
	//
	//	fmt.Println("Loop", i)
	//	for k, v := range m {
	//		fmt.Println(k, v)
	//	}
	//}

	//evenVals := []int{2, 4, 6, 8, 10}
	//for _, v := range evenVals {
	//	v *= 2
	//}
	//fmt.Println(evenVals)

	//
	//words := []string{"a","cow","smile","gopher", "octopus"}
	//for _, word := range words{
	//	switch word{
	//	case "a":
	//		fmt.Println("its a")
	//	case "cow":
	//		fmt.Println("cow")
	//	default:
	//		fmt.Println("Default value")
	//	}
	//}

}
