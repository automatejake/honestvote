package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	string_start := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	start, err := time.Parse(time.RFC1123, string_start)
	if err != nil {
		log.Println(err)
	}

	// a.Format("2006-01-02 15:04:05")

	time.Sleep(1 * time.Second)
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed.Hours() < 4)

}
