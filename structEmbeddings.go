// Struct Embedding Go language ka ek feature hai jo code reuse aur composition ko simple banata hai.
//  Isme ek struct ke andar doosre struct ko include karte hain, jo inheritance jaisa behavior provide karta hai (lekin inheritance se different hai)

package main

import "fmt"

// Without Embeding
// type Address struct {
// 	City    string
// 	ZipCode int
// }

// type Person struct {
// 	Name    string
// 	Age     int
// 	Address Address
// }

//with embedding
type Address struct {
	City    string
	ZipCode int
}

type Person struct {
	Name    string
	Age     int
	Address // Embedded Struct
}

func main() {
	p := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			City:    "Mumbai",
			ZipCode: 400001,
		},
	}
	// fmt.Println("Name:", p.Name)
	// fmt.Println("City:", p.Address.City)
	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.City)

}
