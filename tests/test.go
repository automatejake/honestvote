package main

import (
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	milliseconds_now := time.Now().UnixNano() / 1000000
	fmt.Println(milliseconds_now)
	time.Sleep(time.Second)
	milliseconds_later := time.Now().UnixNano() / 1000000
	fmt.Println(reflect.TypeOf(milliseconds_now))
	fmt.Println(milliseconds_later)

}
