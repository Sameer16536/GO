package main

import "fmt"

func main() {
	nums := []int{1, 2, 33, 4, 4, 5, 5, 6, 6}
	sum := 0
	// _  ignores the index in this case and using range we iterate over values of nums[]
	for _, val := range nums {
		sum = sum + val
	}
	fmt.Println("Sum:", sum)

	for i, val := range nums {
		if val%2 == 0 {
			fmt.Println("index:", i)
		}
	}

	//maps:
	yoMap := map[string]string{"a": "apple", "b": "banana"}
	for i, val := range yoMap {
		fmt.Printf("%s ------> %s\n", i, val)
	}

	for k := range yoMap {
		fmt.Println("Key------->", k)
	}
}
