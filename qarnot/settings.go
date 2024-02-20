package qarnot

import (
	"encoding/json"
	"fmt"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

type Settings struct {
	Storage string `json:"storage"`
}

func (c *Client) GetSettings() (Settings, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, "settings")
	if err != nil {
		return Settings{}, fmt.Errorf("could not get settings due to the following error : %v", err)
	}

	var settings Settings
	err = json.Unmarshal(data, &settings)
	helpers.JsonUnmarshalCheckError(err)

	return settings, nil
}
