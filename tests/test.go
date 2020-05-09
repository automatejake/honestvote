package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// type struct {
// 	ID   int    `json:"id"`
//     Name string `json:"name"`
// }

var databaseDriver string
var databaseUser string
var databaseHost string
var databasePort string
var databaseName string
var eligibleVoterField string

func main() {
	db, err := sql.Open("mysql", "root:test@tcp(registry.honestvote.io:6603)/jacob_is_cool")
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()

	// perform a db.Query insert
	// insert, err := db.Query("CREATE TABLE test ( id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(20) )")
	// insert, err := db.Query("INSERT INTO test VALUES (4, 'TEST' )")
	results, err := db.Query("SELECT name from test")
	for results.Next() {
		var test string
		err = results.Scan(&test)
		fmt.Println(test)
	}

	// // if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer results.Close()
}
