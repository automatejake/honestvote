package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func main() {
	private_key, public_key := crypto.GenerateKeyPair()
	message := []byte("requesting administrator privileges")
	signature, err := crypto.Sign(message, private_key)
	if err != nil {
		return
	}

	var request database.RequestAdminPrivileges = database.RequestAdminPrivileges{
		PublicKey: public_key,
		Signature: signature,
		Message:   message,
	}
	fmt.Println(request)

	json_request, err := json.Marshal(&request)
	if err != nil {
		return
	}

	file, err := os.Create("request-admin-priv.sh")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, _ = io.WriteString(file, "echo \"Election Transaction:\"\n\ncurl --header \"Content-Type: application/json\" --request POST --data '"+
		string(json_request)+"' http://localhost:7003/administrator/request-privileges\n\n\n\n")

}
