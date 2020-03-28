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

func TestGetElectionsHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	database.CollectionPrefix = "a_"
	database.MongoDB = database.MongoConnect("localhost")

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
