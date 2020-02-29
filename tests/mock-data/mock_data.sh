echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Vote for Charity",
	"institutionName": "Honestvote",
	"description": "Whichever charities get the most votes, will be donated $50 each by Honestvote",
	"startDate": "Sat, 29 Feb 2020 16:44:17 EST",
	"endDate": "Wed, 16 Sep 2020 16:44:17 EDT",
	"emailDomain": "^\\w{2}\\d{6}@wcupa\\.edu$",
	"positions": [
		{
			"id": "demfrmeororev",
			"displayName": "Youth Charities",
			"candidates": [
				{
					"name": "Beverlys Birthdays",
					"key": "test1"
				},
				{
					"name": "Brittanys Hope",
					"key": "test2"
				},
				{
					"name": "Ronald McDonald House",
					"key": "test3"
				}
			]
		},
		{
			"id": "defmrfmrkmef",
			"displayName": "Sustainability",
			"candidates": [
				{
					"name": "Art of Recycle",
					"key": "test4"
				},
				{
					"name": "Safe Harbor of Chester County",
					"key": "test5"
				},
				{
					"name": "Global Links",
					"key": "test6"
				}
			]
		}
	],
	"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5",
	"signature": "304602210084818b43a140d99760b39d2ec7fe45de64d370000ba945673048d52b1addb2ad0221009e596631e726564a02d45c2e7763e0b626c49fbf4797261f3cce844dbbaea0e4"
}' http://localhost:7003/election



echo "Registration Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "jacob@neubaum.com",
	"firstName": "Jacob",
	"lastName": "Neubaum",
	"dateOfBirth": "3/9/1999",
	"electionName": "304602210084818b43a140d99760b39d2ec7fe45de64d370000ba945673048d52b1addb2ad0221009e596631e726564a02d45c2e7763e0b626c49fbf4797261f3cce844dbbaea0e4",
	"electionAdmin": "3059301306072a8648ce3d020106082a8648ce3d0301070342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5",
	"publicKey": "3059301306072a8648ce3d020106082a8648ce3d0301070342000433309ab03a630fba855471711ad0ccdb2344479abfd480990228127502a23d447fb2635b876387fa40f91cdab173b34f4b0a7a450afeaa0af59ec7e4d1c24f26",
	"senderSig": "",
	"code": "",
	"timestamp": ""
}' http://localhost:7003/election/test/register



echo "Vote Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Vote",
	"electionId": "304602210084818b43a140d99760b39d2ec7fe45de64d370000ba945673048d52b1addb2ad0221009e596631e726564a02d45c2e7763e0b626c49fbf4797261f3cce844dbbaea0e4",
	"receivers": [
		{
			"positionId": "demfrmeororev",
			"candidateName": "test1"
		}
	],
	"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000433309ab03a630fba855471711ad0ccdb2344479abfd480990228127502a23d447fb2635b876387fa40f91cdab173b34f4b0a7a450afeaa0af59ec7e4d1c24f26",
	"signature": "30460221008eee6e158ff2df002ae05aa2dadd827cf4ecdee4ab015b7808e6b210d104d6c8022100f3403f8d292a60966ca911a0f14c3d1ae499083762aa265873615af785cd0b12"
}' http://localhost:7003/election/test/vote



