package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total = total + num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)

	// we can pass n number of Arguments to the function using Vardiac Function
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	sum(nums...)
}
