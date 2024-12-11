package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Value of i ")
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("Value of j")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	for i := range 3 {
		fmt.Println("range", i)
	}
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
