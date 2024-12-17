package main

import "fmt"

// zeroval has an int parameter, so arguments will be passed to it by value.
//zeroval will get a copy of ival
func zeroVal(ival int) {
	ival = 0
}

//zeroPtr takes in the int  pointer
func zeroPtr(iptr *int) {
	//dereferences the pointer from its memory address to the current
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroVal(i)
	fmt.Println("zeroval:", i)
	// The &i syntax gives the memory address of i, i.e. a pointer to i.

	zeroPtr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	//Simple Example:
	var num int = 42
	fmt.Println("Direct num--->", num)
	//Declaring a pointer to the integer
	var ptr *int
	fmt.Println("Value of ptr -->", ptr) //nil
	ptr = &num
	fmt.Println("Value of ptr after assigning memory address of num to ptr-->", ptr)
	fmt.Println("Dereferencing 'ptr' to get the value---->", *ptr)
	*ptr = 69
	fmt.Println("after changing *ptr value , see the value of num", num)
}
