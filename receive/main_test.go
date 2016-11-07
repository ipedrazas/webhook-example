package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHookHandler(t *testing.T) {

	jsonprep := `{"name":"My Test","password":"Secret Password"}`
	jsonStr := []byte(jsonprep)
	data := bytes.NewBuffer(jsonStr)

	request, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/",
		data,
	)
	if err != nil {
		t.Fatalf("Could not create request , %v", err)
	}
	recorder := httptest.NewRecorder()
	hookHandler(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("Response is not 200 Ok -> %d", recorder.Code)
	}

}
