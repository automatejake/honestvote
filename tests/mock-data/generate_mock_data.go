package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jneubaum/honestvote/tests/mock-data/generatedata"
	"github.com/joho/godotenv"
)

var Election string = "Vote for Charity"
var Institution string = "Honestvote"
var Start string = time.Now().Format(time.RFC1123)
var End string = time.Now().AddDate(0, 200, 200).Format(time.RFC1123) //200 days in  the future
var EmailDomain string = ""
var AdminPriv, AdminPub string
var FileName string = "../tests/mock-data/mock_data.sh"
var ScriptName string = "../scripts/deploy-local-chain.sh"
var Description string = "This is for that"

func main() {
	err := os.Chdir("../../build")
	if err != nil {
		fmt.Println(err)
		return
	}
	godotenv.Load()
	AdminPriv := os.Getenv("PRIVATE_KEY")
	AdminPub := os.Getenv("PUBLIC_KEY")

	generatedata.GenerateMockData(AdminPriv, AdminPub, Election, Institution, Description, Start, End, EmailDomain, FileName, ScriptName)
}
