package main

import (
	"github.com/jneubaum/honestvote/core/core-database/database"
)

type P struct {
	X, Y, Z int
	Name    string
}

func main() {
	database.MongoDB = database.MongoConnect()
	database.ExistsInTable("127.0.0.1", 7002)

}
