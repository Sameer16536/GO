package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)
	//Output : [0 0 0 0 0]

	a[1] = 1
	a[2] = 69
	fmt.Println("Set--->", a)

	fmt.Println("Get-->", a[3])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("B-->", b)

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// If you specify the index with :, the elements in between will be zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoDimArray = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(twoDimArray)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
