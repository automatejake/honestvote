package logger

import (
	"log"
	"os"
)

var Logs bool = false

var f, _ = os.OpenFile("debug.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

var logger = log.New(f, "", log.LstdFlags)

func Println(file_name string, function_name string, message string) {
	log.Println("Filename: " + file_name + "\tFunction name: " + function_name + "\tMessage: " + message)
	logger.Println("Filename: " + file_name + "\tFunction name: " + function_name + "\tMessage: " + message)
}
