package main

import (
	"fmt"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

type Test struct {
	Test  string
	Test2 string
}

func main() {
	database.MongoDB = database.MongoConnect()
	database.CollectionPrefix = "a_"

	s := "7b2252223a36393030393738353332313331313230303333313336353736383737333634363730313137393736373738303039353332323634353133303735303731343132383334373631343330303334332c2253223a393436333037333039353837303139343638323239363431333737363234303734303634383339353938383534303832333731353831323836393432353238373338393536313237353332327d"
	election, err := database.GetElection(s)
	if err != nil {

	}

	//Check to see if election is an ongoing election
	now := time.Now()

	fmt.Println(election)
	electionEnd, err := time.Parse(time.RFC1123, election.End)
	if err != nil {
		fmt.Println(err)
	}
	// time.RFC1123
	if now.After(electionEnd) {
		fmt.Println("hello")
	}
	// priv, pub := crypto.GenerateKeyPair()
	// v := database.Vote{
	// 	Type:     "Vote",
	// 	Election: "BestElection",
	// }
	// encodedV, err := v.Encode()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// hash := crypto.CalculateHash(encodedV)

	// signature, err := crypto.SignTransaction(hash, priv)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// valid, err := crypto.Verify([]byte(hash), database.PublicKey(pub), signature)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// v.Signature = signature
	// fmt.Println(valid)
	// fmt.Println(v)

}
