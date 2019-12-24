package logger

import (
	"log"
	"os"
)

var Logs bool = false

var f, _ = os.OpenFile("debug.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

var logger = log.New(f, "Timestamp: ", log.LstdFlags)

func Println(file_name string, function_name string, message string) {
	if Logs {
		log.Println("\nFilename: " + file_name + "\nFunction name: " + function_name + "\nMessage: " + message + "\n")
		logger.Println("\nFilename: " + file_name + "\nFunction name: " + function_name + "\nMessage: " + message + "\n")
	}
}
