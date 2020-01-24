package main

import (
	"fmt"
	"time"
)

func main() {
	check := time.Time{}
	// start := time.Now().AddDate(1, 1, 1)
	end := time.Now().AddDate(5, 5, 50)
	if check.Before(end) {
		fmt.Println(end)
	}
}
