package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Test  string
	Test2 string
}

func main() {
	var a map[string]interface{}
	fmt.Println(reflect.TypeOf(a).String())
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
