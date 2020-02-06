package main

import (
	"encoding/asn1"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

type Test struct {
}

func main() {
	test := Test{}
	derEncoded, err := asn1.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}

	var test2 database.Vote
	asn1.Unmarshal(derEncoded, &test2)
	fmt.Println(test2)

}
