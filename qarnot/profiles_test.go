package qarnot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestListProfiles(t *testing.T) {
	expected := "[\"guerilla-v2\", \"docker-network\", \"docker-network-ssh\"]"
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/profiles" {
				fmt.Fprint(w, expected)
			}
		}),
	)
	defer srv.Close()

	client, err := NewClient(srv.URL, "xxx", "v1")
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	profiles, _ := client.ListProfiles()
	expectedData := []string{"guerilla-v2", "docker-network", "docker-network-ssh"}

	if !reflect.DeepEqual(profiles, expectedData) {
		t.Errorf("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", profiles)
	}
}

func TestListProfilesBadToken(t *testing.T) {
	expected := "{\"error\":{\"message\":\"Bad authentication token\",\"code\":401},\"data\":{}}"
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(401)
			fmt.Fprint(w, expected)
		}),
	)
	defer srv.Close()

	client, err := NewClient(srv.URL, "xxx", "v1")
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	_, err = client.ListProfiles()
	expectedData := "could not get the list of profiles due to the following error : [HTTP 401] Bad authentication token"

	if err.Error() != expectedData {
		t.Errorf("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetProfileDetails(t *testing.T) {
	expected := `{
		"name": "docker-network",
		"constants": [
		  {
			"name": "DOCKER_SRV",
			"value": "https://registry-1.docker.io",
			"description": "Address of the Docker registry to use, if not the Docker Hub"
		  },
		  {
			"name": "QARNOT_DISABLE_CPU_BOOST",
			"value": "false",
			"description": "Set to 'true' to force-disable the CPU boost. Raw performance will decrease, but can be more predictible over time."
		  }
		],
		"licenses": []
	  }
	`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/profiles/docker-network" {
				fmt.Fprint(w, expected)
			}
		}),
	)
	defer srv.Close()

	client, err := NewClient(srv.URL, "xxx", "v1")
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	profileDetails, _ := client.GetProfileDetails("docker-network")
	expectedData := ProfileDetails{
		Name: "docker-network",
		Constants: []ProfileConstant{
			{
				Name:        "DOCKER_SRV",
				Value:       "https://registry-1.docker.io",
				Description: "Address of the Docker registry to use, if not the Docker Hub",
			},
			{
				Name:        "QARNOT_DISABLE_CPU_BOOST",
				Value:       "false",
				Description: "Set to 'true' to force-disable the CPU boost. Raw performance will decrease, but can be more predictible over time.",
			},
		},
		Licences: nil,
	}

	if !reflect.DeepEqual(profileDetails, expectedData) {
		t.Errorf("different values found")
		t.Errorf("found    : %v", profileDetails)
		t.Errorf("expected : %v", expectedData)
	}
}

func TestGetProfilesDetailsBadToken(t *testing.T) {
	expected := "{\"error\":{\"message\":\"Bad authentication token\",\"code\":401},\"data\":{}}"
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(401)
			fmt.Fprint(w, expected)
		}),
	)
	defer srv.Close()

	client, err := NewClient(srv.URL, "xxx", "v1")
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	_, err = client.GetProfileDetails("test")
	expectedData := "could not get profiles details due to the following error : [HTTP 401] Bad authentication token"

	if err.Error() != expectedData {
		t.Errorf("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", err.Error())
	}
}
