package main

import "fmt"

type Car struct {
	speed int
	model string
}

//We used pointer to directly modify the value of struct
func (c *Car) accelerate() int {
	c.speed = c.speed + 10
	return c.speed
}

func (name Car) sayModel() string {
	return name.model
}

func main() {
	myCar := Car{model: "Hyundai", speed: 60}
	myCar.accelerate()
	fmt.Println("Speed of the car:", myCar.speed)
	myCar.sayModel()
	fmt.Println("Model:", myCar.model)
}
