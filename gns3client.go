package gons3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func get(g GNS3Client, url string, status int, result interface{}) error {
	return req(g, "GET", url, status, nil, result)
}

func delete(g GNS3Client, url string, status int, result interface{}) error {
	return req(g, "DELETE", url, status, nil, result)
}

func post(g GNS3Client, url string, status int, body, result interface{}) error {
	return req(g, "POST", url, status, body, result)
}

func put(g GNS3Client, url string, status int, body, result interface{}) error {
	return req(g, "PUT", url, status, body, result)
}

func req(g GNS3Client, method, url string, status int, body, result interface{}) error {
	// Handle empty body, bytes body, or Marshal body to JSON
	var bodyReader *bytes.Reader
	if bodyBytes, ok := body.([]byte); ok {
		bodyReader = bytes.NewReader(bodyBytes)
	} else if body != nil {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body to json: %w", err)
		}
		bodyReader = bytes.NewReader(reqBody)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	// Create request
	req, err := http.NewRequest(method, g.GetSchemeAuthority()+url, bodyReader)
	defer req.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	// Send request
	resp, err := g.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check status code and return the error if possible
	if resp.StatusCode != status {
		return newServerError(resp)
	}

	// Read body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read http body: %w", err)
	}

	// Unmarshal body as bytes, as JSON, or nothing
	if resultBytes, ok := result.(*[]byte); ok {
		*resultBytes = respBody
	} else if result != nil {
		err = json.Unmarshal(respBody, &result)
		if err != nil {
			return fmt.Errorf("failed to unmarshal http body to json: %v", err)
		}
	}

	return nil
}
