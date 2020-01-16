package http

func HandleFullRoutes() {
	Router.HandleFunc("/elections", GetElectionsHandler).Methods("GET")
	Router.HandleFunc("/election/{electionid}", GetElectionHandler).Methods("GET")
	Router.HandleFunc("/election/{electionid}/candidates", GetCandidatesHandler).Methods("GET")
	Router.HandleFunc("/election/{electionid}/positions", GetPositionsHandler).Methods("GET")
	Router.HandleFunc("/election/{electionid}/tickets", GetTicketsHandler).Methods("GET")
	Router.HandleFunc("/election/{electionid}/votes", GetVotesHandler).Methods("GET")
	Router.HandleFunc("/userpermissions/{publickey}", GetPermissionsHandler).Methods("GET")

	//this needs to be encrypted (send admin's public key and encrypted message containing email and public key)
	Router.HandleFunc("/registerElection/email={email}&public_key={public_key}&election={election}", RegisterHandler).Methods("GET")
	Router.HandleFunc("/registerElection", RegisterHandler).Methods("POST")
}

func HandleProducerRoutes() {
	Router.HandleFunc("/verifyCode/code={id}&verified={verified}&email={email}", VerifyEmailHandler).Methods("GET")
	Router.HandleFunc("/connectpeer", GetPeer).Methods("GET")
	// Should not actually be here, only for testing.  Router.HandleFunc("/test/email={email}&public_key={public_key}&election={election}", RegisterHandler).Methods("GET")
}
