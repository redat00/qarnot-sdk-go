package qarnot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

type Client struct {
	httpClient *http.Client
	url        string
	apiKey     string
	version    string
	s3         *s3.Client
}

// Since the API is not returning consistent errors, we create
// different struct to accomodate them all. If you want more
// information about it, you can check the Github issue :
// https://github.com/redat00/qarnot-sdk-go/issues/7

type firstErrorType struct {
	Message string `json:"message"`
}

func (e *firstErrorType) getErrorString() string {
	return e.Message
}

type secondErrorType struct {
	Error firstErrorType `json:"error"`
}

func (e *secondErrorType) getErrorString() string {
	return e.Error.Message
}

type thirdErrorType struct {
	Errors map[string][]string `json:"errors"`
}

func (e *thirdErrorType) getErrorString() string {
	return fmt.Sprintf("%v", e.Errors)
}

func getErrorStringFromBody(rawMsg []byte) (string, error) {
	var errorString string

	if strings.Contains(string(rawMsg), `"errors": {`) {
		var actualError thirdErrorType
		if err := json.Unmarshal(rawMsg, &actualError); err != nil {
			return "", helpers.FormatJsonUnmarshalError(err)
		}
		errorString = actualError.getErrorString()
	} else if strings.Contains(string(rawMsg), `"error":{`) {
		var actualError secondErrorType
		if err := json.Unmarshal(rawMsg, &actualError); err != nil {
			return "", helpers.FormatJsonUnmarshalError(err)
		}
		errorString = actualError.getErrorString()
	} else {
		var actualError firstErrorType
		if err := json.Unmarshal(rawMsg, &actualError); err != nil {
			return "", helpers.FormatJsonUnmarshalError(err)
		}
		errorString = actualError.getErrorString()
	}

	return errorString, nil
}

func (c *Client) sendRequest(method string, payload []byte, headers map[string]string, endpoint string, options ...func(*http.Request) error) ([]byte, int, error) {
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

	// Apply options
	for _, option := range options {
		option(req)
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
		var rawMsg json.RawMessage
		if err := json.Unmarshal([]byte(body), &rawMsg); err != nil {
			return []byte{}, 0, helpers.FormatJsonUnmarshalError(err)
		}

		errorString, err := getErrorStringFromBody(rawMsg)
		if err != nil {
			panic(err)
		}

		return []byte{}, resp.StatusCode, fmt.Errorf("[HTTP %v] %v", resp.StatusCode, errorString)
	}

	// Return the response
	return body, resp.StatusCode, nil
}

type QarnotConfig struct {
	ApiUrl     string
	ApiKey     string
	Email      string
	Version    string
	StorageUrl string
}

func NewClient(qarnotConfig *QarnotConfig) (*Client, error) {
	// Create an HTTP client
	httpClient := &http.Client{
		Timeout:   15 * time.Second,
		Transport: http.DefaultTransport,
	}

	// Create an AWS Config
	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("default"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(qarnotConfig.Email, qarnotConfig.ApiKey, ""),
		),
	)
	if err != nil {
		return &Client{}, fmt.Errorf("could not create S3 configuration: %v", err)
	}

	// Create an S3 Client
	s3Client := s3.NewFromConfig(awsConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(qarnotConfig.StorageUrl)
	})

	// Create the actual API client
	client := Client{
		httpClient: httpClient,
		url:        qarnotConfig.ApiUrl,
		apiKey:     qarnotConfig.ApiKey,
		version:    qarnotConfig.Version,
		s3:         s3Client,
	}

	// Return the client
	return &client, nil
}
