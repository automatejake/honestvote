echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Student Government Elections",
	"institutionName": "West Chester University",
	"description": "Spring Elections",
	"startDate": "Tue, 11 Feb 2020 19:10:55 EST",
	"endDate": "Sat, 29 Aug 2020 19:10:55 EDT",
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
	"sender": "3059301306072a8648ce3d020106082a8648ce3d030107034200043269053de5c0ee04f06a2aa54087fbe0603a92c1ac4ffd8556ac1687e9f65d48e5c2a0b1b220199a700c1c695e5cf2c2be38f1d042aa4beb770b27e8350cae33",
	"signature": "3045022100e43992dc2b5beb2fc6d606e3ce03b4f7d44fe18fbbe546a951a73bb6ef83a96102201e519f8b99cd2c9b6c303105aa2865f4a93701310e83b69e8fcc07c791cca1da"
}' http://localhost:7003/election



echo "Registration Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "jacob@neubaum.com",
	"firstName": "Jacob",
	"lastName": "Neubaum",
	"dateOfBirth": "3/9/1999",
	"electionName": "3045022100e43992dc2b5beb2fc6d606e3ce03b4f7d44fe18fbbe546a951a73bb6ef83a96102201e519f8b99cd2c9b6c303105aa2865f4a93701310e83b69e8fcc07c791cca1da",
	"electionAdmin": "3059301306072a8648ce3d020106082a8648ce3d030107034200043269053de5c0ee04f06a2aa54087fbe0603a92c1ac4ffd8556ac1687e9f65d48e5c2a0b1b220199a700c1c695e5cf2c2be38f1d042aa4beb770b27e8350cae33",
	"publicKey": "3059301306072a8648ce3d020106082a8648ce3d0301070342000467fb0b123aa28cd362ce69d4ab2946f2689e6009c6d63877698e8545c506654c740b72ab765dc7ecf6f3229ab59a84f90edde2e88e6553c25de2bf5f7d3f1b30",
	"senderSig": "",
	"code": "",
	"timestamp": ""
}' http://localhost:7003/election/test/register



echo "Vote Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Vote",
	"electionName": "3045022100e43992dc2b5beb2fc6d606e3ce03b4f7d44fe18fbbe546a951a73bb6ef83a96102201e519f8b99cd2c9b6c303105aa2865f4a93701310e83b69e8fcc07c791cca1da",
	"receivers": [
		{
			"id": "demfrmeororev",
			"key": "test1"
		}
	],
	"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000467fb0b123aa28cd362ce69d4ab2946f2689e6009c6d63877698e8545c506654c740b72ab765dc7ecf6f3229ab59a84f90edde2e88e6553c25de2bf5f7d3f1b30",
	"signature": "3046022100b447947a790cb7cb89fafe69c8d6c6cf4a3b2dfed1fa12ab527854e791ef81bb022100e43c38baa9915cf704d05549295517434b53208e65a67aa40b05b8f468947986"
}' http://localhost:7003/election/test/vote



