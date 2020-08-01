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

// var MongoConnection = database.MongoConnect(DATABASE_HOST)

func TestPostRequestAdminPrivileges(t *testing.T) {
	// url := "http://127.0.0.1:7003"

	// private_key, public_key := crypto.GenerateKeyPair()
	// message := []byte("requesting administrator privileges")
	// signature, err := crypto.Sign(message, private_key)
	// if err != nil {
	// 	return
	// }

	// var request database.RequestAdminPrivileges = database.RequestAdminPrivileges{
	// 	PublicKey:   public_key,
	// 	Domain:      "bizylife.com",
	// 	Institution: "BizyLife",
	// 	Signature:   signature,
	// 	Message:     message,
	// }

	// json_request, err := json.Marshal(&request)
	// if err != nil {
	// 	return
	// }
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_request))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// req.Header.Set("Content-Type", "application/json")

	// database.CollectionPrefix = CollectionPrefix
	// database.MongoDB = MongoConnection

	// // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(corehttp.PostVoteHandler)

	// // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// // directly and pass in our Request and ResponseRecorder.
	// handler.ServeHTTP(rr, req)

	// // Check the status code is what we expect.
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }
}

func TestGetEndPoint(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	var MongoConnection = database.MongoConnect("localhost")
	database.CollectionPrefix = "a_"
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
	t.Log(endpoint.(string) + " == " + expectedEndpoint)
	if !reflect.DeepEqual(endpoint, expectedEndpoint) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			endpoint, expectedEndpoint)
	}
}
