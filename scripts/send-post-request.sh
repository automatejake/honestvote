curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Student Government Elections",
	"institutionName": "West Chester University",
	"description": "Spring Elections",
	"startDate": "Sat, 01 Feb 2020 19:28:10 EST",
	"endDate": "Wed, 19 Aug 2020 19:28:10 EDT",
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
	"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000417ef2c220ac8c4a69fef5850cc599ef39a9cf2e5ee9298c6bae8a1fc0a8faff93eb2522968faf352e8b96da4f6ea758d61e0c231da525a50dd5fe9f2c63cdec1",
	"signature": "7b2252223a37313735303434303435353636383336373739363834313338323938343532353837363436343131333633363736313633383138353632353731323434393834353832353633383335393635382c2253223a3130323430393934383431313635363634343031363135333430393336363437343032353532393234323733343332373431393538333938353236383930383839343430363436303734303030347d"
}' http://localhost:7003/election