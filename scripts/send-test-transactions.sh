cd "C://Program Files (x86)/Nmap"

echo '{"message":"transaction", "data":{
	"type": "Election",
	"electionName": "Student Government Elections",
	"institutionName": "West Chester University",
	"description": "Spring Elections",
	"startDate": "Fri, 24 Jan 2020 19:53:32 EST",
	"endDate": "Tue, 11 Aug 2020 19:53:32 EDT",
	"emailDomain": "wcupa.edu",
	"positions": [
		{
			"id": "demfrmeororev",
			"displayName": "Student Government President",
			"candidates": [
				{
					"name": "John Doe",
					"key": "test1"
				},
				{
					"name": "Sarah Jennings",
					"key": "test2"
				},
				{
					"name": "Maximus Footless",
					"key": "test3"
				}
			]
		}
	],
	"sender": "3059301306072a8648ce3d020106082a8648ce3d030107034200046a8cee3c7842a24d956d4d4ea10c6255e5f57ec7606e17380cc155781819e9749be050a620e64c2e76152cdea00384cc352f6b28251e8394211173f6038e0f5e",
	"id": "7b2252223a39343931363134343030313031383032393434343933393131333937393334393438323631333234393737303735343331383137333632363634323539343634323435343436353831313235352c2253223a36383834343438323939343837333639303439343234373935323036383934353734383632353933343736343833323431383033353130313333323733373639373032313136393434313338367d"
}, "type":"Election"}' | ./ncat localhost 7004 & \

sleep 1

echo '{"message":"transaction", "data":{"type": "Vote","election": "7b2252223a39343931363134343030313031383032393434343933393131333937393334393438323631333234393737303735343331383137333632363634323539343634323435343436353831313235352c2253223a36383834343438323939343837333639303439343234373935323036383934353734383632353933343736343833323431383033353130313333323733373639373032313136393434313338367d","receiver": {"demfrmeororev": "test1"},"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000447f58a060c071323405b07c7c539a2bbfdc924bb4115895bcd31d19c2389e8296646af03bcc044533d8f8612ed73f15be2e4eec48bcd9a08b1d1b6876598ec39","signature": "7b2252223a37353634393834313038303230363435303933303735373134313137373332303032393734333132383532353232303637393036373137373032353333363738393335353336363433323036332c2253223a393836343733363635303233313636313335363936343536313532393732383930363030333130383233333130393531313634313832353030373330333137323738353039383332313437317d"}, "type":"Vote"}' | ./ncat localhost 7004 & \

sleep 1

echo '{"message":"transaction", "transaction":{
	"type": "Registration",
	"election": "7b2252223a39343931363134343030313031383032393434343933393131333937393334393438323631333234393737303735343331383137333632363634323539343634323435343436353831313235352c2253223a36383834343438323939343837333639303439343234373935323036383934353734383632353933343736343833323431383033353130313333323733373639373032313136393434313338367d",
	"receiver": "test3",
	"sender": "3059301306072a8648ce3d020106082a8648ce3d030107034200046a8cee3c7842a24d956d4d4ea10c6255e5f57ec7606e17380cc155781819e9749be050a620e64c2e76152cdea00384cc352f6b28251e8394211173f6038e0f5e",
	"signature": "7b2252223a31393731323439323935323133303037353938333732313236363030323832303539373435333139343730393733363035313133323130373338373537313930353632353634373332373739302c2253223a32363731353631363532313636383830383732353836353737353036303637303233323030333235303932323539313634323434343337333630373234383636383731313936363132373732377d"
}, "type":"Registration"}' | ./ncat localhost 7004 & \

sleep 20

echo '{"message":"find"}' | ./ncat localhost 7006 & \