package main

import "fmt"

func main() {
	// to define initial values
	m1 := map[string]string{"a": "apple"}

	// to create empty map
	// make also int as second argument to define the initial capacity
	m2 := make(map[string]string)
	m2["a"] = "apple"

	var m3 map[string]string
	// m3["a"] = "apple" // panic error
	m3 = make(map[string]string)
	m3["a"] = "apple"

	m4 := make(map[string]string, 2)
	m4["a"] = "apple"
	m4["b"] = "banana"
	m4["c"] = "cherry"

	fmt.Println("m1", m1)
	fmt.Println("m2", m2)
	fmt.Println("m3", m3)
	fmt.Println("m3", m4)

}
