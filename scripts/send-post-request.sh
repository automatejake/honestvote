curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "akshay.rajaram2000@gmail.com",
	"firstName": "Akshay",
	"lastName": "Rajaram",
	"dateOfBirth": "3/9/1999",
	"electionName": "3046022100d2c0208104bae8cb1b2f2e8ccaa3b13ffe73744bc455d8093633e3de561bdd0d022100eb976b5e3c104a96c0bfdfdd338edd0f494c16574c8474850809486fd6593412",
	"electionAdmin": "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3",
	"publicKey": "03b6a84a459764b044e73a46a615132be0c9360b45458b0828a7ca721cbb41a63c",
	"senderSig": "",
	"code": "",
	"timestamp": "",
	"verified": ""
}' http://localhost:7003/election/test/register

