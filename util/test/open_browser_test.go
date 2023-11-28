package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenbrowser(t *testing.T) {
	// Start a test server to serve a dummy page
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Test opening the URL in a browser
	err := Openbrowser(ts.URL)
	if err != nil {
		t.Errorf("Openbrowser(%s) returned an error: %v", ts.URL, err)
	}
}
