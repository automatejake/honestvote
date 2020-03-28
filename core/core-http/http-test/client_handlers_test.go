package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
	corehttp "github.com/jneubaum/honestvote/core/core-http/http"
)

var CollectionPrefix = "a_"
var MongoConnection = database.MongoConnect("localhost")

func TestGetElectionsHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	database.CollectionPrefix = CollectionPrefix
	database.MongoDB = MongoConnection

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(corehttp.GetElectionsHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	response := corehttp.Payload{}
	_ = json.Unmarshal(rr.Body.Bytes(), &response)

	expected := "OK"
	if response.Status != expected {
		t.Errorf("handler returned unexpected status: got %v want %v",
			response.Status, expected)
	}

	responseData, _ := json.Marshal(&response.Data)
	elections := []database.ElectionInfo{}
	json.Unmarshal(responseData, &elections)

	expectedElections := []database.ElectionInfo{
		database.ElectionInfo{
			ElectionName: "Vote for Charity",
			Institution:  "Honestvote",
			Description:  "Whichever charities get the most votes, will be donated $50 each by Honestvote",
			Start:        "Fri, 27 Feb 2020 08:00:00 EST",
			End:          "Fri, 27 Feb 2020 22:00:00 EST",
			Signature:    "3045022034466fa37fac0368c342705c342bda5e381a8ad92b0209161ca7dc310dfcef8b022100c6249adc9a0d690c4d4b76f32c8fbe226d77e5965d6ed6b234a5733ba76d3504",
		},
	}
	if !reflect.DeepEqual(elections, expectedElections) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseData, expectedElections)
	}

}

// func TestGetElectionHandler(t *testing.T) {
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 	// pass 'nil' as the third parameter.
// 	q := url.Values{}
// 	q.Add("electionid", "3045022034466fa37fac0368c342705c342bda5e381a8ad92b0209161ca7dc310dfcef8b022100c6249adc9a0d690c4d4b76f32c8fbe226d77e5965d6ed6b234a5733ba76d3504")
// 	req, err := http.NewRequest("GET", "/election/tgneij", url.Values{"page": {"1"}, "per_page": {"100"}})
// 	fmt.Println(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	database.CollectionPrefix = CollectionPrefix
// 	database.MongoDB = MongoConnection

// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(corehttp.GetElectionHandler)

// 	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// 	// directly and pass in our Request and ResponseRecorder.
// 	handler.ServeHTTP(rr, req)

// 	// Check the status code is what we expect.
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	// Check the response body is what we expect.
// 	response := corehttp.Payload{}
// 	_ = json.Unmarshal(rr.Body.Bytes(), &response)

// 	expected := "OK"
// 	if response.Status != expected {
// 		t.Errorf("handler returned unexpected status: got %v want %v",
// 			response.Status, expected)
// 	}

// 	responseData, _ := json.Marshal(&response.Data)
// 	election := database.Election{}
// 	json.Unmarshal(responseData, &election)

// 	expectedElection := database.Election{
// 		Type:         "Election",
// 		ElectionName: "Vote for Charity",
// 		Institution:  "Honestvote",
// 		Description:  "Whichever charities get the most votes, will be donated $50 each by Honestvote",
// 		Start:        "Fri, 27 Feb 2020 08:00:00 EST",
// 		End:          "Fri, 27 Feb 2020 22:00:00 EST",
// 		EmailDomain:  "^\\w{2}\\d{6}@wcupa\\.edu$",
// 		Positions:    []database.Position{},
// 		Sender:       "3059301306072a8648ce3d020106082a8648ce3d03010703420004f84767e25c0d12886513c7309f5ff2c636ee61b71d9c4934cc50c22d83eda89b8e3a4e2279d5181dd69627b9bf6ef688d5888fe8e3b17b6a0f94944adf7d34b1",
// 		Signature:    "3045022034466fa37fac0368c342705c342bda5e381a8ad92b0209161ca7dc310dfcef8b022100c6249adc9a0d690c4d4b76f32c8fbe226d77e5965d6ed6b234a5733ba76d3504",
// 	}

// 	if !reflect.DeepEqual(election, expectedElection) {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			election, expectedElection)
// 	}

// }
