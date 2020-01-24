package main

import "fmt"

func main() {
	m := make(map[string]bool)
	m["hi"] = true
	fmt.Println(m["hi"])
	fmt.Println(m["hiw"])
}
