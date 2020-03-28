package main

import (
	"testing"
)

func TestGetEndPoint(t *testing.T) {
	// // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// // pass 'nil' as the third parameter.
	// req, err := http.NewRequest("GET", "/health-check", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// database.CollectionPrefix = CollectionPrefix
	// database.MongoDB = MongoConnection

	// // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(corehttp.GetElectionsHandler)

	// // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// // directly and pass in our Request and ResponseRecorder.
	// handler.ServeHTTP(rr, req)

	// // Check the status code is what we expect.
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

	// // Check the response body is what we expect.
	// response := corehttp.Payload{}
	// _ = json.Unmarshal(rr.Body.Bytes(), &response)

	// expected := "OK"
	// if response.Status != expected {
	// 	t.Errorf("handler returned unexpected status: got %v want %v",
	// 		response.Status, expected)
	// }

	// responseData, _ := json.Marshal(&response.Data)
	// elections := []database.ElectionInfo{}
	// json.Unmarshal(responseData, &elections)

	// expectedEn
	// if !reflect.DeepEqual(elections, expectedElections) {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		responseData, expectedElections)
	// }

}
