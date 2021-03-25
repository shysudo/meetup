package update

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateParticipantHandler(t *testing.T) {
	var jsonStr = []byte(`{"id":11,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`)

	req, err := http.NewRequest("PUT", "/meetup/participants", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateParticipantHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":11,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
