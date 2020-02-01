curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Student Government Elections",
	"institutionName": "West Chester University",
	"description": "Spring Elections",
	"startDate": "Sat, 01 Feb 2020 14:27:27 EST",
	"endDate": "Wed, 19 Aug 2020 14:27:27 EDT",
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
	"sender": "3059301306072a8648ce3d020106082a8648ce3d030107034200047c1d5cd365f84677317b8aa975421ad08d1242b66d5163534e9e20f814e40a14e464f300658aaf0ff5e1551bf805290ba6c7ca4ee3b0b0c8abf920fe51861ca6",
	"signature": "7b2252223a31323438363239303436343733343938313136373238313537343935363936393034353338383436353239323732333131373239393330323534343331363534333530363738313734373236312c2253223a33373031373830383331333631393333323731333034373635323039393235343532353537323530373039383633323734383034363536373935363233343037333138313031343430363139347d"
}' http://localhost:7003/election