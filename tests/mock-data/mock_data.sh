echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Vote for Charity",
	"institutionName": "West Chester University",
	"description": "Whichever charities get the most votes, will be donated $50 each by Honestvote",
	"startDate": "Wed, 26 Feb 2020 13:18:15 EST",
	"endDate": "Sun, 13 Sep 2020 13:18:15 EDT",
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
	"signature": "3046022100ed94b1c26cbc9084370f5c8fc42705381f4dfd52678b566cf13dd4c800425876022100876b06affe21813d918d003aa33a9ecaa346f6304ba603d47e1f29d72e401674"
}' http://localhost:7003/election



echo "Registration Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "jacob@neubaum.com",
	"firstName": "Jacob",
	"lastName": "Neubaum",
	"dateOfBirth": "3/9/1999",
	"electionName": "3046022100ed94b1c26cbc9084370f5c8fc42705381f4dfd52678b566cf13dd4c800425876022100876b06affe21813d918d003aa33a9ecaa346f6304ba603d47e1f29d72e401674",
	"electionAdmin": "3059301306072a8648ce3d020106082a8648ce3d0301070342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5",
	"publicKey": "3059301306072a8648ce3d020106082a8648ce3d030107034200043870783f641130b7fedd13ad176effb966fd0d4789f32bd2749abf1f18dae7ff3f48bd581c598b219220f57b58c523404b16dea124ea97ce3a7ad25cd93c8893",
	"senderSig": "",
	"code": "",
	"timestamp": ""
}' http://localhost:7003/election/test/register



echo "Vote Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Vote",
	"electionId": "3046022100ed94b1c26cbc9084370f5c8fc42705381f4dfd52678b566cf13dd4c800425876022100876b06affe21813d918d003aa33a9ecaa346f6304ba603d47e1f29d72e401674",
	"receivers": [
		{
			"id": "demfrmeororev",
			"key": "test1"
		}
	],
	"sender": "3059301306072a8648ce3d020106082a8648ce3d030107034200043870783f641130b7fedd13ad176effb966fd0d4789f32bd2749abf1f18dae7ff3f48bd581c598b219220f57b58c523404b16dea124ea97ce3a7ad25cd93c8893",
	"signature": "3045022100ca03e723b5d31f0ca92f38b806b5afbb70b8a5ab079514d5953e4a3aae659c15022008f1c23095c79a0bfd1334526c49c687fc649aaddef1df72d2157bcc8cf9224a"
}' http://localhost:7003/election/test/vote



