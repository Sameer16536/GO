// Enums ka main purpose hai related constant values ko ek group mein organize karna.
// Enums ek tarah ke constant values ka set hota hai, jo kisi specific category ko represent karta hai.
// Go mein enums ko achieve karne ke liye const aur iota ka use karte hain.
package main

import "fmt"

const (
	Sunday = iota //iota --> automatically starts from 0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//iota ek special keyword hai jo har new constant group mein automatically 0 se start hota hai aur 1 se increment hota hai.

//Enums with custom values
const (
	Small  = 1
	Medium = 5
	Large  = 10
)

func main() {
	fmt.Println("Sunday:", Sunday)       // Output: Sunday: 0
	fmt.Println("Monday:", Monday)       // Output: Monday: 1
	fmt.Println("Tuesday:", Tuesday)     // Output: Tuesday: 2
	fmt.Println("Wednesday:", Wednesday) // Output: Wednesday: 3
	fmt.Println("Thursday:", Thursday)   // Output: Thursday: 4
	fmt.Println("Friday:", Friday)       // Output: Friday: 5
	fmt.Println("Saturday:", Saturday)   // Output: Saturday: 6

	fmt.Println("Small size:", Small)   // Output: Small size: 1
	fmt.Println("Medium size:", Medium) // Output: Medium size: 5
	fmt.Println("Large size:", Large)   // Output: Large size: 10
}
