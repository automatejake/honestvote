package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var okay map[string]int

func main() {
	godotenv.Load()
	a := os.Getenv("TEST")

	err := os.Chdir("../../build")
	if err != nil {
		fmt.Println(err)
	}
	godotenv.Load()

	b := os.Getenv("KEY")
	fmt.Println(a, b)
}
