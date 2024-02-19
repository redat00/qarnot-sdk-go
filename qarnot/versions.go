package qarnot

import "encoding/json"

type Version struct {
	Version   string
	EndOfLife string
}

func (c *Client) GetVersions() []Version {
	// Send request and get back data
	data, _ := c.sendRequest("GET", []byte{}, make(map[string]string), "versions")

	// Convert data to UserInfo struct
	var versions []Version
	err := json.Unmarshal(data, &versions)
	if err != nil {
		panic(err)
	}

	// Return Versions
	return versions
}
