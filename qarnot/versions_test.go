package qarnot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestGetVersions(t *testing.T) {
	expected := "[{\"version\": \"v0.1\",\"endOfLife\": \"2020-03-16\"}]"
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/versions" {
				fmt.Fprint(w, expected)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	versions, _ := client.GetVersions()
	expectedData := Version{Version: "v0.1", EndOfLife: "2020-03-16"}
	if !slices.Contains(versions, expectedData) {
		t.Errorf("error in values. Expected %+v, found %+v", expectedData, versions)
	}
}
