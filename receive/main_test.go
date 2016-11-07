package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHookHandler(t *testing.T) {

	cases := []struct {
		in   string
		code int
	}{
		{`{"name":"My Test","password":"Secret Password"}`, http.StatusOK},
		{`{"name":"My Test","password": }`, http.StatusUnprocessableEntity},
		{"", http.StatusUnprocessableEntity},
	}

	for _, c := range cases {
		jsonStr := []byte(c.in)
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
		if recorder.Code != c.code {
			t.Errorf("Response is not 200 Ok -> %d", recorder.Code)
		}
	}

}
