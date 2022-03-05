package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//import "fmt"

//type person struct {
//	age  int
//	name string
//}
//
//func main() {
//	p := person{}
//	i := 2
//	s := "Hello"
//	tryToModify(i, s, p)
//	fmt.Println(i, s, p)
//}
//
//func tryToModify(i int, s string, p person) {
//	i = i * 2
//	s = "Goodbye"
//	p.name = "Zeytin"
//}

//-------------- //

//func main() {
//	m := map[int]string{
//		1: "first",
//		2: "second",
//	}
//	modMap(m)
//	fmt.Println(m)
//
//	s := []int{1, 2, 3}
//	modSlice(s)
//	fmt.Println(s)
//}
//
//func modMap(m map[int]string) {
//	m[2] = "hello"
//	m[3] = "goodbye"
//	delete(m, 1)
//}
//
//func modSlice(s []int) {
//	for k, v := range s {
//		s[k] = v * 2
//	}
//	s = append(s, 10)
//}

//-------POINTERS--------//
//type person struct {
//	age  int
//	name string
//}
//
//func main() {
//	p := person{}
//	i := 2
//	s := "Hello"
//	tryToModify(&i, &s, &p)
//	fmt.Println(i, s, p)
//}
//
//func tryToModify(i *int, s *string, p *person) {
//	text := "Goodbye"
//	*i = *i * 2
//	*s = text
//	p.name = "Zeytin"
//}

//------------//
//
//type person struct {
//	FirstName  string
//	MiddleName *string
//	LastName   string
//}

//func main() {
//p := person{
//	FirstName:  "Pat",
//	MiddleName: "Perry", // This line won't compile
//	LastName:   "Peterson",
//}

//p := person{
//	FirstName:  "Pat",
//	MiddleName: stringp("Perry"), // This works. ps: use a helper function to turn to constant value into a pointer
//	LastName:   "Peterson",
//}

//p := person{
//	FirstName: "Pat",
//	LastName:  "Peterson",
//}
//
//fmt.Println(p.FirstName, p.MiddleName, p.LastName)
//
//}

//func stringp(s string) *string {
//	return &s
//}

//------ DO NOT FEAR THE POINTERS------//
//func main() {
//	var f *int
//	willItUpdate(f)
//	fmt.Println(f)
//}
//
//func willItUpdate(g *int) {
//	x := 8
//	g = &x
//}

//-----------//

type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
}

//dont do this
//func MakeFoo(p *Person) error {
//	p.FirstName = "mustafa"
//	p.MiddleName = "kemal"
//	p.LastName = "atatürk"
//	return nil
//}
//
////do this
//func MakePerson(p Person) (Person, error) {
//	p.FirstName = "mustafa"
//	p.MiddleName = "kemal"
//	p.LastName = "atatürk"
//	return p, nil
//}

//-------------//
func main() {
	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	err := json.Unmarshal([]byte(`{"name": "Mustafa", "age": 30}`), &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f)
}
