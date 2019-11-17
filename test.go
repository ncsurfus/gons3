package gons3

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func newTestClient(do func(req *http.Request) (*http.Response, error)) testGNS3Client {
	return testGNS3Client{do: do}
}

type testGNS3Client struct {
	do func(req *http.Request) (*http.Response, error)
}

func (t testGNS3Client) GetSchemeAuthority() string {
	return ""
}

func (t testGNS3Client) Do(req *http.Request) (*http.Response, error) {
	return t.do(req)
}

func (t testGNS3Client) CreateResponse(status int, body interface{}) *http.Response {
	var response = &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
	}

	// Handle empty body, bytes body, or Marshal body to JSON
	switch b := body.(type) {
	case nil:
		response.Body = ioutil.NopCloser(bytes.NewReader([]byte{}))
	case *[]byte:
		response.Body = ioutil.NopCloser(bytes.NewReader(*b))
		response.Header.Add("Content-Type", "application/octet-stream")
	default:
		reqBody, _ := json.Marshal(body)
		response.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
		response.Header.Add("Content-Type", "application/json")
	}

	return response
}

func createTestBody(b interface{}) io.ReadCloser {
	if b == nil {
		return ioutil.NopCloser(bytes.NewReader([]byte{}))
	}
	j, _ := json.Marshal(b)
	return ioutil.NopCloser(bytes.NewReader(j))
}
