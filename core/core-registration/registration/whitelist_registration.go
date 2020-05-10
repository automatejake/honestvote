package registration

import (
	"github.com/jneubaum/honestvote/core/core-database/database"

	_ "github.com/go-sql-driver/mysql"
)

/*
* 	The purpose of this function is for election administrators to specify a sql data source to verify whether or not a someone is eligible to vote
 */
func OnWhitelist(Registrant string, whitelist database.WhiteListElectionSettings) bool {
	// db, err := sql.Open(whitelist.DatabaseDriver, whitelist.DatabaseUser+":"+whitelist.DatabaseName+"@tcp("+whitelist.DatabaseHost+":"+whitelist.DatabasePort+")/"+whitelist.TableName)
	// if err != nil {
	// 	logger.Println("whitelist_registration", "OnWhiteList", err)
	// }
	// defer db.Close()

	// results, err := db.Query("SELECT " + whitelist.EligibleVoterField + " FROM " + whitelist.TableName + " WHERE " + whitelist.EligibleVoterField + " = " + Registrant)
	// var voter string
	// for results.Next() {
	// 	err = results.Scan(&voter)
	// 	if voter == Registrant {
	// 		logger.Println("whitelist_registration", "OnWhiteList", "Registrant exists on whitelist")
	// 		return true
	// 	}

	// }
	// if err != nil {
	// 	logger.Println("whitelist_registration", "OnWhiteList", err)
	// }

	// defer results.Close()

	// return false
	return true
}
