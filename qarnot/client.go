package qarnot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

type Client struct {
	httpClient http.Client
	url        string
	apiKey     string
	version    string
}

type subErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type errorResponse struct {
	Message string           `json:"message,omitempty"`
	Error   subErrorResponse `json:"error,omitempty"`
}

func (c *Client) sendRequest(method string, payload []byte, headers map[string]string, endpoint string) ([]byte, int, error) {
	// Build the request using url and endpoint
	var req *http.Request
	var err error
	if len(payload) > 0 {
		req, err = http.NewRequest(method, fmt.Sprintf("%v/%v/%v", c.url, c.version, endpoint), bytes.NewReader(payload))
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%v/%v/%v", c.url, c.version, endpoint), nil)
	}
	if err != nil {
		return []byte{}, 0, fmt.Errorf("could not create request due to the following error: %v", err)
	}

	// Add required headers to request
	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	// Add more headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

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
		var reqError errorResponse
		err := json.Unmarshal(body, &reqError)
		helpers.JsonUnmarshalCheckError(err)

		var message string
		if reqError.Message != "" {
			message = reqError.Message
		} else {
			message = reqError.Error.Message
		}

		return []byte{}, resp.StatusCode, fmt.Errorf("[HTTP %v] %v", resp.StatusCode, message)
	}

	// Return the response
	return body, resp.StatusCode, nil
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
