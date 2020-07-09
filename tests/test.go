package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	now := time.Now().Format(time.RFC1123)

	fmt.Println(now)

	parsley, _ := time.Parse(now, ".Format(time.RFC1123)")
	fmt.Println(parsley.Before(time.Now()))
}
