package qarnot

import "encoding/json"

type Settings struct {
	Storage string `json:"storage"`
}

func (c *Client) GetSettings() Settings {
	data, _ := c.sendRequest("GET", []byte{}, nil, "settings")

	var settings Settings
	err := json.Unmarshal(data, &settings)
	if err != nil {
		panic(err)
	}

	return settings
}
