curl --header "Content-Type: application/json" --request POST --data '{
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
}' http://localhost:7003/election