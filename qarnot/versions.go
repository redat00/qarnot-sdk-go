package qarnot

import (
	"encoding/json"
	"fmt"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

func (c *Client) GetVersions() ([]Version, error) {
	data, _, err := c.sendRequest("GET", []byte{}, make(map[string]string), "versions")
	if err != nil {
		return []Version{}, fmt.Errorf("could not get versions due to the following error : %v", err)
	}

	var versions []Version
	err = json.Unmarshal(data, &versions)
	helpers.JsonUnmarshalCheckError(err)

	return versions, nil
}

type Version struct {
	Version   string
	EndOfLife string
}
