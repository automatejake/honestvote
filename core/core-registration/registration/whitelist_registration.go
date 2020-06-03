package registration

import (
	"database/sql"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

/*
* 	The purpose of this function is for election administrators to specify a sql data source to verify whether or not a someone is eligible to vote
 */
func OnWhitelist(Registrant string, whitelist database.WhiteListElectionSettings) bool {
	logger.Println("whitelist_registration", "OnWhiteList", whitelist.DatabaseUser+":"+whitelist.DatabasePassword+"@tcp("+whitelist.DatabaseHost+":"+whitelist.DatabasePort+")/"+whitelist.DatabaseName)
	db, err := sql.Open(whitelist.DatabaseDriver, whitelist.DatabaseUser+":"+whitelist.DatabasePassword+"@tcp("+whitelist.DatabaseHost+":"+whitelist.DatabasePort+")/"+whitelist.DatabaseName)
	if err != nil {
		logger.Println("whitelist_registration", "OnWhiteList", err)
	}
	defer db.Close()

	results, err := db.Query("SELECT " + whitelist.EligibleVoterField + " FROM " + whitelist.TableName + " WHERE " + whitelist.EligibleVoterField + " = '" + Registrant + "'")
	var voter string
	for results.Next() {

		err = results.Scan(&voter)
		logger.Println("whitelist_registration", "OnWhiteList", voter)
		if voter == Registrant {
			logger.Println("whitelist_registration", "OnWhiteList", "Registrant exists on whitelist")
			return true
		}

	}
	if err != nil {
		logger.Println("whitelist_registration", "OnWhiteList", err)
	}

	defer results.Close()
	logger.Println("whitelist_registration", "OnWhiteList", "Registrant does not exist on whitelist")
	return false
}
