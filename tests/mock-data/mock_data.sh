echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Student Government Elections",
	"institutionName": "West Chester University",
	"description": "Spring Elections",
	"startDate": "Sun, 16 Feb 2020 10:24:11 EST",
	"endDate": "Thu, 03 Sep 2020 10:24:11 EDT",
	"emailDomain": "^\\w{2}\\d{6}@wcupa\\.edu$",
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
	"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5",
	"signature": "3045022100cd48a8cdc1f1178532cfe2390847f0edcc549cc8173b1c59be8876a0826d57760220645a019ebb29c60920d24942390406ac9a05c84085e70d7b730976fbd7556f43"
}' http://localhost:7003/election



echo "Registration Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "jacob@neubaum.com",
	"firstName": "Jacob",
	"lastName": "Neubaum",
	"dateOfBirth": "3/9/1999",
	"electionName": "3045022100cd48a8cdc1f1178532cfe2390847f0edcc549cc8173b1c59be8876a0826d57760220645a019ebb29c60920d24942390406ac9a05c84085e70d7b730976fbd7556f43",
	"electionAdmin": "3059301306072a8648ce3d020106082a8648ce3d0301070342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5",
	"publicKey": "3059301306072a8648ce3d020106082a8648ce3d03010703420004be78e22479cdf3e71430e58858574b350ab9a425163994a2ff604bdc9a0433a4cfe4d30008354fefa0b12defcb7018da965680de972958ffdc63e5612bd5e910",
	"senderSig": "",
	"code": "",
	"timestamp": ""
}' http://localhost:7003/election/test/register



echo "Vote Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Vote",
	"electionId": "3045022100cd48a8cdc1f1178532cfe2390847f0edcc549cc8173b1c59be8876a0826d57760220645a019ebb29c60920d24942390406ac9a05c84085e70d7b730976fbd7556f43",
	"receivers": [
		{
			"id": "demfrmeororev",
			"key": "test1"
		}
	],
	"sender": "3059301306072a8648ce3d020106082a8648ce3d03010703420004be78e22479cdf3e71430e58858574b350ab9a425163994a2ff604bdc9a0433a4cfe4d30008354fefa0b12defcb7018da965680de972958ffdc63e5612bd5e910",
	"signature": "304502206879d7d8d11285d579dd65775fe91287a424bc381e6b2181122feb3ec100915d022100ca553c5a142530ea00ba5d415e9f0248e29e2018aa708dc50fea0609f18224f4"
}' http://localhost:7003/election/test/vote



