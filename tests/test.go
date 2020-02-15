package main

import "fmt"

var Test map[string]int = map[string]int{"test": 1}

func main() {

	for a, b := range Test {
		fmt.Println("a is: ", a)
		fmt.Println("b is: ", b)
	}
}
