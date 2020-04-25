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

func Println(file_name string, function_name string, message interface{}) {
	message_string := fmt.Sprintf("%v", message)

	switch Mode {
	case "All":
		log.Println("\nFilename: " + file_name + "\nFunction name: " + function_name + "\nMessage: " + message_string + "\n")
		logger.Println("\nFilename: " + file_name + "\nFunction name: " + function_name + "\nMessage: " + message_string + "\n\n")
	case "Debug":
		fmt.Println("\nFunction name:" + function_name + "\n" + message_string)
		logger.Println("\nFunction name:" + function_name + "\n" + message_string)
	case "Info":
		fmt.Println(message_string)
		logger.Println(message_string)
	}

}
