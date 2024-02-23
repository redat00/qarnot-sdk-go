package qarnot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

type subErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type errorResponse struct {
	Message string           `json:"message,omitempty"`
	Error   subErrorResponse `json:"error,omitempty"`
}

type Client struct {
	httpClient *http.Client
	url        string
	apiKey     string
	version    string
	s3         *s3.Client
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
		if err != nil {
			return []byte{}, 0, helpers.FormatJsonUnmarshalError(err)
		}

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
