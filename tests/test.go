package main

import (
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func main() {
	// vote := database.Vote{
	// 	Type:     "Vote",
	// 	Election: "election1",
	// 	Receiver: []database.SelectedCandidate{
	// 		database.SelectedCandidate{
	// 			PositionId: "position2",
	// 			Recipient:  "Trisha Newton",
	// 		},
	// 		database.SelectedCandidate{
	// 			PositionId: "position3",
	// 			Recipient:  "Casey Brady",
	// 		},
	// 	},
	// 	Sender:    "30819e134d3635363937353032313735353739303638333638343636393931313635383538373638353337363630363932343039363438383234303939363631333137343234353234323833383939333134134d3833363432353133373436383435333239393831313333323533373539303237373335333938323536323534393531373239333332333035373430363630353231393332303231393938343839",
	// 	Signature: "3044022005c4efa3bfde40486bbe3b352a8873ed2abd852b1dd46a56bd163d13cd7e8c8102206ae5bab7d4443187f09ce3a82fcc22648e41a07fa157481879af1c8696b6eb73",
	// }

	// _, err := validation.IsValidVote(vote)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// ifaces, err := net.Interfaces()
	// if err != nil {

	// }
	// // handle err
	// for _, i := range ifaces {
	// 	addrs, err := i.Addrs()
	// 	if err != nil {

	// 	}
	// 	// handle err
	// 	for _, addr := range addrs {
	// 		var ip net.IP
	// 		switch v := addr.(type) {
	// 		case *net.IPNet:
	// 			ip = v.IP
	// 		case *net.IPAddr:
	// 			ip = v.IP
	// 		}
	// 		// process IP address
	// 		fmt.Println(ip)
	// 	}
	// }
	database.CollectionPrefix = "a_"
	database.MongoDB = database.MongoConnect("localhost")
	// fmt.Println(database.coon

}
