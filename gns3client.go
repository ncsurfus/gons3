package gons3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GNS3HTTPClient represents a default GNS3 client and server.
type GNS3HTTPClient struct {
	Client   *http.Client
	Scheme   string
	Hostname string
	Port     int
}

// GetSchemeAuthority gets the scheme and authority of the GNS3 server.
func (g GNS3HTTPClient) GetSchemeAuthority() string {
	if g.Scheme == "" {
		g.Scheme = "http"
	}
	if g.Hostname == "" {
		g.Hostname = "127.0.0.1"
	}
	if g.Port == 0 {
		g.Port = 3080
	}
	return fmt.Sprintf("%v://%v:%v", g.Scheme, g.Hostname, g.Port)
}

// Do sends the HTTP request with the default or explicit *http.Client.
func (g GNS3HTTPClient) Do(req *http.Request) (*http.Response, error) {
	if g.Client == nil {
		return http.DefaultClient.Do(req)
	}
	return g.Client.Do(req)
}

// GNS3Client provides an interface for creating custom GNS3 clients.
type GNS3Client interface {
	GetSchemeAuthority() string
	Do(req *http.Request) (*http.Response, error)
}

func get(g GNS3Client, url string, expectedStatus int, result interface{}) error {
	return req(g, "GET", url, expectedStatus, nil, result)
}

func delete(g GNS3Client, url string, expectedStatus int, result interface{}) error {
	return req(g, "DELETE", url, expectedStatus, nil, result)
}

func post(g GNS3Client, url string, expectedStatus int, body, result interface{}) error {
	return req(g, "POST", url, expectedStatus, body, result)
}

func put(g GNS3Client, url string, expectedStatus int, body, result interface{}) error {
	return req(g, "PUT", url, expectedStatus, body, result)
}

// ErrFailedToMarshalBodyToJSON is returned when the body could not be marshaled to JSON.
var ErrFailedToMarshalBodyToJSON = errors.New("failed to marshal body to json")

// ErrFailedToCreateRequest is returned when the request could not be created.
var ErrFailedToCreateRequest = errors.New("failed to create request")

// ErrRequestFailed is returned when the request failed.
var ErrRequestFailed = errors.New("request failed")

// ErrUnexpectedStatusCode is returned when the status code indicates an error.
var ErrUnexpectedStatusCode = errors.New("unexpected status code")

// ErrFailedToReadResult is returned when response result could not be read into a buffer.
var ErrFailedToReadResult = errors.New("failed to read response body")

// ErrResponseNotJSON is returned when response was expected to be JSON, but was not.
var ErrResponseNotJSON = errors.New("response was not json as expected")

// ErrFailedToUnmarshalResponse is returned when the json response could not be unmarshled.
var ErrFailedToUnmarshalResponse = errors.New("failed to unmarshal response")

func req(g GNS3Client, method, url string, expectedStatus int, body, result interface{}) error {
	// Handle empty body, bytes body, or Marshal body to JSON
	var bodyReader *bytes.Reader
	var contentType string
	switch b := body.(type) {
	case nil:
		bodyReader = bytes.NewReader([]byte{})
	case []byte:
		bodyReader = bytes.NewReader(b)
		contentType = "application/octet-stream"
	default:
		reqBody, err := json.Marshal(body)
		if err != nil {
			return Wrap(ErrFailedToMarshalBodyToJSON, err)
		}
		bodyReader = bytes.NewReader(reqBody)
		contentType = "application/json"
	}

	// Create request
	req, err := http.NewRequest(method, g.GetSchemeAuthority()+url, bodyReader)
	if err != nil {
		return Wrap(ErrFailedToCreateRequest, err)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	defer req.Body.Close()

	// Send request
	resp, err := g.Do(req)
	if err != nil {
		return Wrap(ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	// Check status code and return the error if possible
	if resp.StatusCode != expectedStatus {
		return Wrap(ErrUnexpectedStatusCode, newServerError(resp))
	}

	// Read body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Wrap(ErrFailedToReadResult, err)
	}

	// Read response
	switch r := result.(type) {
	case nil:
	case *[]byte:
		*r = respBody
	default:
		if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
			return ErrResponseNotJSON
		}
		if json.Unmarshal(respBody, &result); err != nil {
			return Wrap(ErrFailedToUnmarshalResponse, err)
		}
	}

	return nil
}
