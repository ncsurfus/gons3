package gons3

import (
	"bytes"
	"encoding/json"
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
		g.Client = http.DefaultClient
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
			return fmt.Errorf("failed to marshal body to json: %w", err)
		}
		bodyReader = bytes.NewReader(reqBody)
		contentType = "application/json"
	}

	// Create request
	req, err := http.NewRequest(method, g.GetSchemeAuthority()+url, bodyReader)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	defer req.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Send request
	resp, err := g.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check status code and return the error if possible
	if resp.StatusCode != expectedStatus {
		return newServerError(resp)
	}

	// Read body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read repsonse body: %w", err)
	}

	// Read response
	switch r := result.(type) {
	case nil:
	case *[]byte:
		*r = respBody
	default:
		if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
			return fmt.Errorf("response was not JSON as expected")
		}
		if json.Unmarshal(respBody, &result); err != nil {
			return fmt.Errorf("failed to unmarshal result to json: %v", err)
		}
	}

	return nil
}
