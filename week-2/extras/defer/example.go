// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	B()
}

func A() {
	defer D()
	fmt.Print("P")
}

func B() {
	C()
	fmt.Print("E")
}

func C() {
	defer E()
	defer A()
}

func D() {
	fmt.Print("A")
}

func E() {
	fmt.Print("C")
}

func F() {
	fmt.Print("B")
}
