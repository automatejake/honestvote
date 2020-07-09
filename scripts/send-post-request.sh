curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "akshay.rajaram2000@gmail.com",
	"firstName": "Jacob",
	"lastName": "Neubaum",
	"dateOfBirth": "3/9/1999",
	"electionName": "304502204efeca15f4b0134f96b2d1788dd96bb7c280484c93306e204a0b8450abb1b33e02210095153d021f8b5e6ee3d5947de652cbce9a5b045a9edf3744f5b12d5ebd2f3a9e",
	"electionAdmin": "03e4973aa880e3dbd88210a063a3878ec80d5a688083f60f63577c3d6f0db784c0",
	"publicKey": "037dae85174fd41d9cb891e25b099b022b0bfca66895db8ac65ef3358c1de44f57",
	"senderSig": "",
	"code": "",
	"timestamp": "",
	"verified": ""
}' http://localhost:7003/election/test/register

