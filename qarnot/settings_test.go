package qarnot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSettingsOK(t *testing.T) {
	expected := "{\"storage\": \"https://storage.qarnot.com\"}"
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/settings" {
				fmt.Fprint(w, expected)
			}
		}),
	)
	defer srv.Close()

	client, err := NewClient(srv.URL, "xxx", "v1")
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	settings, _ := client.GetSettings()
	expectedData := Settings{Storage: "https://storage.qarnot.com"}

	if settings != expectedData {
		t.Errorf("different value.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", settings)
	}
}
