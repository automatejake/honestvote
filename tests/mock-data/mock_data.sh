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
<<<<<<< HEAD
	"sender": "3059301306072a8648ce3d020106082a8648ce3d030107034200041ca9ef0f92d2a1143cea968205160e3ebb24dad7b63feef36a231445201cd42d89a2303b24819266d53848db9b60bff507560f7df362c0fd82ca1fc19408656f",
	"signature": "3044022023cb8a231336885ff2d1f32b8909ba031ce3fe8bc65f664b5f89714834659fdf0220027a30509d3e478a90f7718f22429b577df6fc5a3f4462ece03903674e22d42e"
=======
	"sender": "3059301306072a8648ce3d020106082a8648ce3d0301070342000420f6ae9be26dfde8b50f550bfb273ad77d1012a9c427f4e5ea761faa108ab0b69a042448b15e09c67075cba02931c2ae602b9125afad8f0480f83d24c55d3bc5",
	"signature": "3045022100cd48a8cdc1f1178532cfe2390847f0edcc549cc8173b1c59be8876a0826d57760220645a019ebb29c60920d24942390406ac9a05c84085e70d7b730976fbd7556f43"
<<<<<<< HEAD
>>>>>>> 0c7e579f74e62359838040d6660050c5fc8ec257
}' http://localhost:7003/election
=======
}' http://portainer.honestvote.io:7003/election
>>>>>>> 50d3e5012a2b075a55befe1d7391cbbd3761fac8



# echo "Registration Transaction:"

<<<<<<< HEAD
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
=======
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
>>>>>>> 0c7e579f74e62359838040d6660050c5fc8ec257



# echo "Vote Transaction:"

<<<<<<< HEAD
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
=======
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
>>>>>>> 0c7e579f74e62359838040d6660050c5fc8ec257



