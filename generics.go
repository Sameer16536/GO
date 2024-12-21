// Generics Go programming language ka ek feature hai jo code ko type-safe aur reusable banata hai.
// Generics aapko ek single function ya data structure likhne ki facility dete hain jo multiple types ke saath kaam kar sakta hai
// Why Generics?:::::::
// Reusability: Har type ke liye alag-alag function likhne ki zarurat nahi hoti.
// Type Safety: Type-casting se bachata hai aur compile-time par errors detect karta hai.
// Code Duplication Avoidance: Ek hi generic function multiple data types ke saath kaam kar sakta hai.

package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println("Stack:::", intStack)
	fmt.Println("Popped:", intStack.Pop()) // Output: 20

	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")
	fmt.Println("Popped:", stringStack.Pop()) // Output: World
}
