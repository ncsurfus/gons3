package gons3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Transport defines a set of methods to interact with a GNS3 server.
type Transport interface {
	Get(path string, result interface{}) (*http.Response, error)
	Delete(path string, result interface{}) (*http.Response, error)
	Post(path string, body, result interface{}) (*http.Response, error)
	Put(path string, body, result interface{}) (*http.Response, error)
}

// GNS3HTTPServer implements a set of methods to interact with GNS3 via HTTP.
type GNS3HTTPServer struct {
	Hostname string
}

// Get sends an HTTP GET request to the GNS3 server.
func (t GNS3HTTPServer) Get(path string, result interface{}) (*http.Response, error) {
	return t.req("GET", path, nil, result)
}

// Delete sends an HTTP DELETE request to the GNS3 server.
func (t GNS3HTTPServer) Delete(path string, result interface{}) (*http.Response, error) {
	return t.req("DELETE", path, nil, result)
}

// Post sends an HTTP POST request to the GNS3 server.
func (t GNS3HTTPServer) Post(path string, body, result interface{}) (*http.Response, error) {
	return t.req("POST", path, body, result)
}

// Put sends an HTTP PUT request to the GNS3 server.
func (t GNS3HTTPServer) Put(path string, body interface{}, result interface{}) (*http.Response, error) {
	return t.req("PUT", path, body, result)
}

func (t GNS3HTTPServer) req(method, path string, body, result interface{}) (*http.Response, error) {
	// Handle empty body, bytes body, or Marshal body to JSON
	var bodyReader *bytes.Reader
	if body == nil {
		bodyReader = bytes.NewReader([]byte{})
	} else if bodyBytes, ok := body.([]byte); ok {
		bodyReader = bytes.NewReader(bodyBytes)
	} else {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body to json: %v", err)
		}
		bodyReader = bytes.NewReader(reqBody)
	}

	// Create request
	url := "http://" + t.Hostname + path
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %v", err)
	}

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp, fmt.Errorf("http request failed: %v", err)
	}

	// Check Status
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("http request failed: %v", resp.Status)
	}

	// Read body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf("failed to read http body: %v", err)
	}

	if result != nil {
		// Unmarshal body to JSON
		err = json.Unmarshal(respBody, &result)
		if err != nil {
			return resp, fmt.Errorf("failed to unmarshal http body to json: %v", err)
		}
	}

	return resp, nil
}
