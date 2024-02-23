package qarnot

import (
	"encoding/json"
	"fmt"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

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

func (c *Client) ListProfiles() ([]string, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, "profiles")
	if err != nil {
		return []string{}, fmt.Errorf("could not get the list of profiles due to the following error : %v", err)
	}

	var profiles []string
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		return nil, helpers.FormatJsonUnmarshalError(err)
	}

	return profiles, nil
}

func (c *Client) GetProfileDetails(name string) (ProfileDetails, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("profiles/%v", name))
	if err != nil {
		return ProfileDetails{}, fmt.Errorf("could not get profiles details due to the following error : %v", err)
	}

	var profileDetails ProfileDetails
	err = json.Unmarshal(data, &profileDetails)
	if err != nil {
		return ProfileDetails{}, helpers.FormatJsonUnmarshalError(err)
	}

	return profileDetails, nil
}
