package main

import (
	"strings"
)

type Car struct {
	Brand string
	Model string
	Year  int
	Color string
	Price int
}

func main() {
	car := NewCar()
	car = car.SetPrice(200)

	//car2 := &Car{Price: 200, Model: strings.ToLower("Renault CLİO")}
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) SetPrice(price int) *Car {
	c.Price = 200
	return c
}

func (c *Car) SetModel(model string) *Car {
	c.Model = strings.ToLower(model)
	return c
}
