package gons3

import (
	"net/http"
	"testing"
)

func TestGetNode(t *testing.T) {
	// Arrange
	c, req, body := testGNS3Client{}, &http.Request{}, Node{Name: "TestName"}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(200, body), nil
	}

	// Act
	n, err := GetNode(c, "P1", "N1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/N1" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "GET" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
	if n.Name != body.Name {
		t.Errorf("Name is invalid: %v", n.Name)
	}
}

func TestGetNodes(t *testing.T) {
	// Arrange
	c, req := testGNS3Client{}, &http.Request{}
	body := []Node{Node{Name: "TestName1"}, Node{Name: "TestName2"}}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(200, body), nil
	}

	// Act
	n, err := GetNodes(c, "P1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/N1" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "GET" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
	if len(n) != len(body) {
		t.Fatalf("Nodes length is invalid: %v", len(n))
	}
	if n[0].Name != body[0].Name {
		t.Errorf("Name[0] is invalid: %v", n[0].Name)
	}
	if n[1].Name != body[1].Name {
		t.Errorf("Name[1] is invalid: %v", n[1].Name)
	}
}

func TestSuspendNode(t *testing.T) {
	// Arrange
	c, req, body := testGNS3Client{}, &http.Request{}, Node{Name: "TestName"}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(200, body), nil
	}

	// Act
	n, err := SuspendNode(c, "P1", "N1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/N1/suspend" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "POST" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
	if n.Name != body.Name {
		t.Errorf("Name is invalid: %v", n.Name)
	}
}

func TestSuspendNodes(t *testing.T) {
	// Arrange
	c, req := testGNS3Client{}, &http.Request{}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(204, nil), nil
	}
	// Act
	err := SuspendNodes(c, "P1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/suspend" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "POST" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
}

func TestStopNode(t *testing.T) {
	// Arrange
	c, req, body := testGNS3Client{}, &http.Request{}, Node{Name: "TestName"}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(200, body), nil
	}
	// Act
	n, err := StopNode(c, "P1", "N1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/N1/stop" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "POST" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
	if n.Name != body.Name {
		t.Errorf("Name is invalid: %v", n.Name)
	}
}

func TestStopNodes(t *testing.T) {
	// Arrange
	c, req := testGNS3Client{}, &http.Request{}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(204, nil), nil
	}
	// Act
	err := StopNodes(c, "P1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/stop" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "POST" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
}

func TestReloadNode(t *testing.T) {
	// Arrange
	c, req, body := testGNS3Client{}, &http.Request{}, Node{Name: "TestName"}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(200, body), nil
	}
	// Act
	n, err := ReloadNode(c, "P1", "N1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/N1/reload" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "POST" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
	if n.Name != body.Name {
		t.Errorf("Name is invalid: %v", n.Name)
	}
}

func TestReloadNodes(t *testing.T) {
	// Arrange
	c, req := testGNS3Client{}, &http.Request{}
	c.do = func(r *http.Request) (*http.Response, error) {
		req = r
		return c.CreateResponse(204, nil), nil
	}
	// Act
	err := ReloadNodes(c, "P1")

	// Assert
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}
	if req.URL.String() != "/v2/projects/P1/nodes/reload" {
		t.Errorf("URL is invalid: %v", req.URL)
	}
	if req.Method != "POST" {
		t.Errorf("Method is invalid: %v", req.Method)
	}
}
