package main

import "fmt"

type Car struct {
	engineState bool
}

type CarIR interface {
	StartEngine()
	StopEngine()
}

func (c *Car) StartEngine() {
	c.engineState = true
}

func (c *Car) StopEngine() {
	c.engineState = false
}
func main() {
	var car Car
	car.StartEngine()
	fmt.Println(car.engineState)
	car.StopEngine()
	fmt.Println(car.engineState)
}
