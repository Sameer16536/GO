package main

import "fmt"

func closure() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {
	counter := closure()
	fmt.Println("Value this time of x", counter())
	fmt.Println("Value this time of x", counter())
	fmt.Println("Value this time of x", counter())
	fmt.Println("Value this time of x", counter())

	nums := []int{1, 2, 3, 4, 5, 6}

	// Define a closure to filter even numbers
	isEven := func(n int) bool {
		return n%2 == 0
	}

	evens := filter(nums, isEven)
	fmt.Println(evens) // Output: [2 4 6]

}

func filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}
