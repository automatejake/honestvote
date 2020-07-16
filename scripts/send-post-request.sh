curl --header "Content-Type: application/json" --request POST --data '{
	"emailAddress": "akshay.rajaram2000@gmail.com",
	"firstName": "Jacob",
	"lastName": "Neubaum",
	"dateOfBirth": "3/9/1999",
	"electionName": "3045022100c55f0057a8c7751efd22abbba2552aa2baff4f019c91dfe9d63f7a03e927c584022056f3cce733b9fafaa3606bf184b43e1e764d0df9665921fcc4c599fd224aecce",
	"electionAdmin": "03e4973aa880e3dbd88210a063a3878ec80d5a688083f60f63577c3d6f0db784c0",
	"publicKey": "02ba7da6823bb84de2d99b6452ded809ee82798f8453867d9c929eec9bac42cb0d",
	"senderSig": "",
	"code": "",
	"timestamp": "",
	"verified": ""
}' http://localhost:7003/election/test/register

