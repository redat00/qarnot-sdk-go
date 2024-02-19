package qarnot

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListProfiles() []string {
	data, _ := c.sendRequest("GET", []byte{}, nil, "profiles")

	var profiles []string
	err := json.Unmarshal(data, &profiles)
	if err != nil {
		panic(err)
	}

	return profiles
}

type ProfileLicence struct {
	Name         string
	MaxInstances int
	MaxCores     int
}

type ProfileConstant struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type ProfileDetails struct {
	Name      string            `json:"name"`
	Constants []ProfileConstant `json:"constants"`
	Licences  []ProfileLicence  `json:"licences"`
}

func (c *Client) GetProfileDetails(name string) ProfileDetails {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("profiles/%v", name))

	var profileDetails ProfileDetails
	err := json.Unmarshal(data, &profileDetails)
	if err != nil {
		panic(err)
	}

	return profileDetails
}
