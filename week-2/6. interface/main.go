package main

import (
	"fmt"
)

//func main() {
//	var i interface{}
//	i = 20
//	i = "hello"
//	i = struct {
//		FirstName string
//		LastName  string
//	}{"Mustafa", "Kemal"}
//
//	fmt.Println(i)
//}

//---------//

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsFemale bool   `json:"isFemale"`
	Details  struct {
		FirstKey  string `json:"firstKey"`
		SecondKey string `json:"secondKey"`
	} `json:"details"`
}

//
//func main() {
//data := map[string]interface{}{}
//contents, err := ioutil.ReadFile("data.json")
//if err != nil {
//	panic(err)
//}
//
//if err := json.Unmarshal(contents, &data); err != nil {
//	panic(err)
//}

//fmt.Println(data)

//fmt.Println(data["age"].(float64) + 1)
//
//fmt.Println(len(data["name"].(string)))
//
//fmt.Println(!data["isFemale"].(bool))
//
//details := data["details"].(map[string]interface{})
//for _, d := range details {
//	fmt.Println(d)
//}

//data := Person{}
//contents, err := ioutil.ReadFile("data.json")
//if err != nil {
//	panic(err)
//}
//
//if err := json.Unmarshal(contents, &data); err != nil {
//	panic(err)
//}
//
////fmt.Println(data.IsFemale,data.Name,data.Details.SecondKey)
//
//byteData, err := json.Marshal(data)
//if err != nil {
//	panic(err)
//}
//
//fmt.Println(byteData)
//fmt.Println(string(byteData))
//}

type Animal interface {
	Sound() interface{}
	Eat(bool) string
}

type Dog struct {
}

func (d Dog) Sound() interface{} {
	return 123
}

func (d Dog) Eat(b bool) string {
	return "dog-eat!"
}

type Cat struct {
}

func (c Cat) Sound() interface{} {
	return "Miyav!"
}
func (c Cat) Eat(b bool) string {
	return "cat-eat!"
}

type Cow struct {
}

func (l Cow) Sound() interface{} {
	return "Moo!"
}
func (l Cow) Eat(b bool) string {
	return "cow-eat!"
}

// if you got method you are my type
func main() {
	animals := []Animal{Dog{}, Cat{}, Cow{}}

	//dog:=Dog{}
	//dog.Sound()
	//
	//cat := Cat{}
	//cat.Sound()

	for _, animal := range animals {
		fmt.Println(animal.Sound())
		fmt.Println(animal.Eat(true))
	}
}
