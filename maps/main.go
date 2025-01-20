package main

import "fmt"

func main() {
	m1 := map[string]string{"a": "apple"}

	m2 := make(map[string]string)
	m2["a"] = "apple"

	var m3 map[string]string
	// m3["a"] = "apple" // panic error
	m3 = make(map[string]string)
	m3["a"] = "apple"

	fmt.Println("m1", m1)
	fmt.Println("m2", m2)
	fmt.Println("m3", m3)

}
