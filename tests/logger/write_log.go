package logger

import (
	"fmt"
	"log"
	"os"
)

var Mode string = "Debug"

var f, _ = os.OpenFile("debug.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

var logger = log.New(f, "Timestamp: ", log.LstdFlags)

func Println(file_name string, function_name string, message string) {
	switch Mode {
	case "All":
		log.Println("\nFilename: " + file_name + "\nFunction name: " + function_name + "\nMessage: " + message + "\n")
		logger.Println("\nFilename: " + file_name + "\nFunction name: " + function_name + "\nMessage: " + message + "\n")
	case "Debug":
		fmt.Println("\nFunction name:" + function_name + "\n" + message)
		logger.Println("\nFunction name:" + function_name + "\n" + message)
	case "Info":
		fmt.Println(message)
		logger.Println(message)
	}

}
