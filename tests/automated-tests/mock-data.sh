echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Vote for Charity",
	"institutionName": "Honestvote",
	"description": "Whichever charities get the most votes, will be donated $50 each by Honestvote",
	"startDate": "Sat, 25 Apr 2020 12:50:05 EDT",
	"endDate": "Mon, 13 Jul 2037 12:50:05 EDT",
	"emailDomain": "",
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
	"sender": "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3",
	"signature": "3046022100bbb4e8ed9694d7ea6ebb40fb48b2b3cf8f861a6979ad36c1bf6d40a71585068f022100bb23a10bf55f0bf866baef193a016839d1d23add1a19b7db651c429a02e0baf5"
}' http://localhost:7003/election



echo "Registration Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "jacob@neubaum.com",
	"firstName": "Jacob",
	"lastName": "Neubaum",
	"dateOfBirth": "3/9/1999",
	"electionName": "3046022100bbb4e8ed9694d7ea6ebb40fb48b2b3cf8f861a6979ad36c1bf6d40a71585068f022100bb23a10bf55f0bf866baef193a016839d1d23add1a19b7db651c429a02e0baf5",
	"electionAdmin": "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3",
	"publicKey": "037706bb5168fda9337b7445f275f4de9c7e87fbb8cb3f2429ff9a08718d4aa568",
	"senderSig": "",
	"code": "",
	"timestamp": ""
}' http://localhost:7003/election/test/register



echo "Vote Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Vote",
	"electionId": "3046022100bbb4e8ed9694d7ea6ebb40fb48b2b3cf8f861a6979ad36c1bf6d40a71585068f022100bb23a10bf55f0bf866baef193a016839d1d23add1a19b7db651c429a02e0baf5",
	"receivers": [
		{
			"positionId": "demfrmeororev",
			"candidateName": "Beverlys Birthdays"
		}
	],
	"sender": "037706bb5168fda9337b7445f275f4de9c7e87fbb8cb3f2429ff9a08718d4aa568",
	"signature": "30460221008d864bf9498e51119a19763cba41d1790bdfc18451c81c53050a9b8cf042c216022100b870fbe8a00ff286f80c80cc4ef3ee1737e7264172f2dd0ae479616259f4c001"
}' http://localhost:7003/election/test/vote



