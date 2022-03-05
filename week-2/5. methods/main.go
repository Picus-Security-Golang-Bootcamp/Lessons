package main

import "fmt"

//type Person struct {
//	FirstName string
//	LastName  string
//	Age       int
//	TestF     func() bool
//}
//
////func (struct) functionName() string
//func (p Person) String() string {
//	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
//}
//
//func (p Person) ChangeAge() int {
//	return p.Age
//}
//
//func main() {
//	person := Person{
//		FirstName: "Mustafa",
//		LastName:  "Kemal",
//		Age:       30,
//	}
//	fmt.Println(person.String())
//	fmt.Println(person.ChangeAge())
//
//	person.TestF = func() bool {
//		return false
//	}
//}

//--------------//

//type Adder struct {
//	start int
//}
//
//func (a Adder) AddTo(val int) int {
//	return a.start + val
//}
//
//func main() {
//	myAdder := Adder{start: 10}
//	fmt.Println(myAdder.AddTo(5))
//}

//--------------//
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	return []Employee{}
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{
			{Name: "Ahmet", ID: "123"},
			{Name: "Fatma", ID: "9659"},
		},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())
}

//
//type Inner struct {
//	ID int
//}
//
//type Outer struct {
//	Inner
//	ID int
//}
//
//func main() {
//	inner := Inner{
//		ID: 10,
//	}
//
//	o := Outer{
//		Inner: inner,
//		ID:    20,
//	}
//
//	fmt.Println(o.ID)
//	fmt.Println(o.Inner.ID)
//}
