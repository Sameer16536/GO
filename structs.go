package main

import "fmt"

// Structs allow you to encapsulate these fields into a single entity. /// i.e Collection : Grouping same category entities

//lowercase name of type means its private
//Uppercase public
type person struct {
	name string
	age  int
}
type Employee struct {
	name      string
	age       int
	isWorking bool
}

// You can create instances of a struct in several ways:
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	employee1 := Employee{
		name:      "Sameer",
		age:       21,
		isWorking: true,
	}
	//Accessing the data
	fmt.Println("Employee 1 name:", employee1.name)
	fmt.Println("Employee 1 age:", employee1.age)

	person1 := person{
		name: "XYZ",
		age:  19,
	}

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("John boi"))

	fmt.Printf("person name is %s and age is %s\n", person1.name, person1.age)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)

	//Annonymous struct
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
