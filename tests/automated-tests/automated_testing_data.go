package main

import (
	"github.com/jneubaum/honestvote/tests/mock-data/generatedata"
)

var AdminPriv string = "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676"
var AdminPub string = "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"
var Election string = "Vote for Charity"
var Institution string = "Honestvote"
var Description string = "Whichever charities get the most votes, will be donated $50 each by Honestvote"
var Start string = "Sat, 25 Apr 2020 12:50:05 EDT"
var End string = "Mon, 13 Jul 2037 12:50:05 EDT"
var EmailDomain string = ""
var FileName string = "mock-data.sh"
var ScriptName string = "startup-script.sh"

func main() {
	generatedata.GenerateMockData(AdminPriv, AdminPub, Election, Institution, Description, Start, End, EmailDomain, FileName, ScriptName)
}
