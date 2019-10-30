package main_test

import (
	"."
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize()
	code := m.Run()

	os.Exit(code)
}

func TestEmptyEvents(t *testing.T) {

	req, _ := http.NewRequest("GET", "/events", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentEvent(t *testing.T) {

	req, _ := http.NewRequest("GET", "/event/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Event not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Event not found'. Got '%s'", m["error"])
	}
}

func TestCreateEvent(t *testing.T) {

	payload := []byte(`{"name":  "test event","price": 11.22}`)

	req, _ := http.NewRequest("POST", "/event", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test event" {
		t.Errorf("Expected event name to be 'test event'. Got '%v'", m["name"])
	}

	if m["price"] != 11.22 {
		t.Errorf("Expected event price to be '11.22'. Got '%v'", m["price"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected event ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetEvent(t *testing.T) {
	addEvents(1)

	req, _ := http.NewRequest("GET", "/event/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addEvents(count int) {
	if count < 1 {
		count = 1
	}

	// todo
}

func TestUpdateEvent(t *testing.T) {
	addEvents(1)

	req, _ := http.NewRequest("GET", "/event/1", nil)
	response := executeRequest(req)
	var originalEvent map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalEvent)

	payload := []byte(`{"name":"test event - updated name","price":11.22}`)

	req, _ = http.NewRequest("PUT", "/event/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalEvent["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalEvent["id"], m["id"])
	}

	if m["name"] == originalEvent["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalEvent["name"], m["name"], m["name"])
	}

	if m["price"] == originalEvent["price"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", originalEvent["price"], m["price"], m["price"])
	}
}

func TestDeleteEvent(t *testing.T) {
	addEvents(1)

	req, _ := http.NewRequest("GET", "/event/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/event/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/event/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
