package qarnot

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	url        string
	apiKey     string
	version    string
}

func (c *Client) sendRequest(method string, payload []byte, headers map[string]string, endpoint string) ([]byte, int) {
	// Build the request using url and endpoint
	var req *http.Request
	var err error
	if len(payload) > 0 {
		req, err = http.NewRequest(method, fmt.Sprintf("%v/%v/%v", c.url, c.version, endpoint), bytes.NewReader(payload))
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%v/%v/%v", c.url, c.version, endpoint), nil)
	}
	if err != nil {
		return []byte{}, 0
	}

	// Add required headers to request
	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	// Launch the request using the HTTP client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		panic(fmt.Errorf("an error happened during the execution of the request: %v", err))
	}
	defer resp.Body.Close()

	// Read the content of the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Check that the request did not fail
	if resp.StatusCode >= 400 {
		panic(fmt.Errorf("an error happened during the treatment of the request : \nendpoint: %v\nbody:%v", endpoint, string(body)))
	}

	// Return the response
	return body, resp.StatusCode
}

func NewClient(url string, apiKey string, version string) (*Client, error) {
	// Create an HTTP client
	httpClient := &http.Client{
		Timeout:   15 * time.Second,
		Transport: http.DefaultTransport,
	}

	// Create the actual API client
	client := Client{
		httpClient: *httpClient,
		url:        url,
		apiKey:     apiKey,
		version:    version,
	}

	// Return the client
	return &client, nil
}
