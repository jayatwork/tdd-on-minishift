//testing our handlers
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TesthandleRequests(t *testing.T) {

	ctx := MakeMyContext()
	req, err := http.NewRequest("GET", "/" , nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(http.HandlerFunc(handleRequests))

	handler(ctx, rr,req)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned bad status code: received %v and wanted %v", status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: received %v but needed %v", rr.Body.String(), expected)
	}
}

