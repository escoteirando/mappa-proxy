package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMappa(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/hc", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}
