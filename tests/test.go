package main

import (
	"fmt"
	"time"
)

func main() {
	// blah := time.Now()
	// start := time.Now().String()
	// // Thu, 01 Dec 1994 16:00:00 GMT

	// end, err := time.Parse("2006-01-02T15:04:05.000Z", start)
	// if err != nil {
	// 	logger.Println("email_registration.go", "IsValidRegistrationCode()", err.Error())
	// }

	// fmt.Println(blah, "\n", start, "\n", end)

	// var HOURS float64 = 4
	// if time.Now().Sub(end).Hours() > HOURS {
	// 	fmt.Println("We didnt make it")
	// } else {
	// 	fmt.Println("We made it")
	// }
	// time.RFC3339
	blah := time.Now().Format(time.RFC3339) //"Dec 29, 2014 at 7:54pm (SGT)"
	start, _ := time.Parse(time.RFC3339, blah)
	end := start.Add(4 * time.Hour)
	now := time.Now()
	// time.Sleep(3 * time.Second)

	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(now)
	if end.Before(now.Add()) {
		fmt.Println("made it here")
	} else {
		fmt.Println("made it fear")
	}

	// vote := database.Vote{
	// 	Type:     "Vote",
	// 	Election: "Chester",
	// 	Receiver: map[string]string{"cool": "beans"},
	// }

	// voteHeaders := vote.Type + vote.Election
	// for key, value := range vote.Receiver {
	// 	voteHeaders += key + value
	// }

	// fmt.Println([]byte(voteHeaders))

}
