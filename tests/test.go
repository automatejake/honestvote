package main

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func main() {
	database.MongoDB = database.MongoConnect()
	database.CollectionPrefix = "a_"
	fmt.Println(database.FindNode("tests").Role == "producer")

}
