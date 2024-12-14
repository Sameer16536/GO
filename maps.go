package main

import (
	"fmt"
	"maps"
)

func main() {
	// key is string and value is Int
	var m = make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2
	fmt.Println("Map-->", m)

	//Access the value of the key
	v := m["key1"]
	fmt.Println("Value of V -->", v)

	//Accessing the key which is not Present will return 0
	v1 := m["key3"]
	fmt.Println("value of key which doesnt exist", v1)

	//length of the map
	fmt.Println("Length -->", len(m))

	delete(m, "key2")
	fmt.Println("Deleted key 2 from m ", m)

	//Remove all Key-value pairs
	clear(m)
	fmt.Println("cleared M ", m)

	n := map[string]int{"One": 1, "two": 2}
	fmt.Println("new map", n)

	n2 := map[string]int{"One": 1, "two": 2, "three": 3}
	if maps.Equal(n, n2) {
		fmt.Println("Both are equal")
	} else {
		fmt.Println("Not equal")
	}
}
