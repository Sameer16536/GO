package main

import "fmt"

func twoSum(a int, b int) int {
	ans := (a + b)
	return ans
}

func mult(a int, b int) int {
	return a * b
}

//Multiple Return values
//eg return both result and error

func vals() (int, int) {
	return 3, 9
}

func finalResult(sumResult int, multResult int) (string, string) {
	var sumMessage, multMessage string

	// Conditional message for sum
	if sumResult < 1 {
		sumMessage = "Mast hai"
	} else {
		sumMessage = "Bakwass"
	}

	// Conditional message for multiply
	if multResult > 1 {
		multMessage = "Mast hai"
	} else {
		multMessage = "Bakwass"
	}

	return sumMessage, multMessage
}

func main() {
	//Call the functions here in main
	result := twoSum(1, 2)
	fmt.Println(result)

	multilpy := mult(69, 1)
	fmt.Println("multilpy", multilpy)

	a, b := vals()
	fmt.Println("A", a)
	fmt.Println("B", b)

	sumMessage, multMessage := finalResult(result, multilpy)
	fmt.Println("Sum Message:", sumMessage)
	fmt.Println("Multiply Message:", multMessage)

}
