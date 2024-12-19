package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Interface ek blueprint hota hai jo define karta hai ki ek struct (ya type) ko kaunse methods implement karne chahiye.
// Yeh ensure karta hai ki different structs ek common behavior follow karen.

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type
type Circle struct {
	Radius float64
}

// Rectangle type
type Rectangle struct {
	Width, Height float64
}

// Implementing Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implementing Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Function to print details of any Shape
func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

// Example 2
type IndianCricketTeam interface {
	runsScored() int
}

type RohitSharma struct {
	captaincy int
	strokes   int
	century   int
	trophies  int
}

func (r *RohitSharma) runsScored() int {
	runs := r.captaincy + r.century + r.strokes + r.trophies
	return runs

}

type ViratKohli struct {
	captaincy      int
	strokes        int
	century        int
	trophies       int
	aggressiveness int
}

func (v *ViratKohli) runsScored() int {
	runs := v.captaincy + v.century + v.strokes + v.trophies*v.aggressiveness
	return runs
}

func main() {
	// Create instances of Circle and Rectangle
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	// Both Circle and Rectangle implement the Shape interface
	printShapeDetails(c)
	printShapeDetails(r)

	// Example 2
	rohit := RohitSharma{
		captaincy: rand.Intn(10),
		strokes:   rand.Intn(10),
		century:   rand.Intn(10),
		trophies:  rand.Intn(10),
	}
	fmt.Println("Runs Scored by Rohit Sharma:", rohit.runsScored())

	virat := ViratKohli{
		captaincy:      rand.Intn(10),
		strokes:        rand.Intn(10),
		century:        rand.Intn(10),
		trophies:       rand.Intn(10),
		aggressiveness: rand.Intn(100),
	}
	fmt.Println("Runs Scored by Virat Kohli:", virat.runsScored())

}

//Benefits of interface:
// Common Behavior: Interface ensure karta hai ki alag-alag structs ek jaise behavior follow karen.
// Example: runsScored() method dono structs me hai, par unka implementation alag ho sakta hai.
// Polymorphism: Interface ka use karke hum ek generic code likh sakte hain jo different types ke objects ke saath kaam kare.
// Flexibility: Interface tightly coupled nahi hota, isiliye ek struct ko asani se modify kiya ja sakta hai bina baaki code ko todhe.
