package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	test := []int{1, 2, 3, 4, 5, 6}
	test = append(test, 7)
	fmt.Println(test[0])
	test = test[1:]

}
