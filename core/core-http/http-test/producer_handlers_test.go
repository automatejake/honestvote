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

func TestGetEndPoint(t *testing.T) {
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
	handler := http.HandlerFunc(corehttp.GetEndpoint)

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

	endpoint := response.Data

	expectedEndpoint := "127.0.0.1:7002"
	if !reflect.DeepEqual(endpoint, expectedEndpoint) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			endpoint, expectedEndpoint)
	}
}
