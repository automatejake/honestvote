echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{
	"type": "Election",
	"electionName": "Student Government Elections",
	"institutionName": "West Chester University",
	"description": "Spring Elections",
	"startDate": "Sat, 15 Feb 2020 11:59:42 EST",
	"endDate": "Wed, 02 Sep 2020 11:59:42 EDT",
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
	"sender": "3059301306072a8648ce3d020106082a8648ce3d030107034200041ca9ef0f92d2a1143cea968205160e3ebb24dad7b63feef36a231445201cd42d89a2303b24819266d53848db9b60bff507560f7df362c0fd82ca1fc19408656f",
	"signature": "3044022023cb8a231336885ff2d1f32b8909ba031ce3fe8bc65f664b5f89714834659fdf0220027a30509d3e478a90f7718f22429b577df6fc5a3f4462ece03903674e22d42e"
}' http://localhost:7003/election



# echo "Registration Transaction:"

# curl --header "Content-Type: application/json" --request POST --data '{
# 	"emailAddress": "jacob@neubaum.com",
# 	"firstName": "Jacob",
# 	"lastName": "Neubaum",
# 	"dateOfBirth": "3/9/1999",
# 	"electionName": "3044022023cb8a231336885ff2d1f32b8909ba031ce3fe8bc65f664b5f89714834659fdf0220027a30509d3e478a90f7718f22429b577df6fc5a3f4462ece03903674e22d42e",
# 	"electionAdmin": "3059301306072a8648ce3d020106082a8648ce3d030107034200041ca9ef0f92d2a1143cea968205160e3ebb24dad7b63feef36a231445201cd42d89a2303b24819266d53848db9b60bff507560f7df362c0fd82ca1fc19408656f",
# 	"publicKey": "3059301306072a8648ce3d020106082a8648ce3d0301070342000489688124184460028a05bdd5b1800695cd912ae45e5c939a462e30c6dc4802678eeed9efb65a8bb365a9dcc5a3a65e65e8fc346f0dc5eb1267c205e5ad033e77",
# 	"senderSig": "",
# 	"code": "",
# 	"timestamp": ""
# }' http://localhost:7003/election/test/register



# echo "Vote Transaction:"

# curl --header "Content-Type: application/json" --request POST --data '{
# 	"type": "Vote",
# 	"electionName": "3044022023cb8a231336885ff2d1f32b8909ba031ce3fe8bc65f664b5f89714834659fdf0220027a30509d3e478a90f7718f22429b577df6fc5a3f4462ece03903674e22d42e",
# 	"receivers": [
# 		{
# 			"id": "demfrmeororev",
# 			"key": "test1"
# 		}
# 	],
# 	"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000489688124184460028a05bdd5b1800695cd912ae45e5c939a462e30c6dc4802678eeed9efb65a8bb365a9dcc5a3a65e65e8fc346f0dc5eb1267c205e5ad033e77",
# 	"signature": "3045022079a1cb2b6c4a7d5e6eb39e885d292485e7d81b2e3b84edb875467841b8ef9724022100d82a41a46a2cf8ac814707a996f5ceca1eb8760ae99ead62d273a22fe9521046"
# }' http://localhost:7003/election/test/vote



