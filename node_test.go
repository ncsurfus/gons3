package gons3

import (
	"net/http"
	"testing"
)

func TestSuspend(t *testing.T) {
	c, count, nbody := testGNS3Client{}, 0, Node{Name: "TestName"}
	c.do = func(req *http.Request) (*http.Response, error) {
		count++
		if req.URL.String() != "/v2/projects/P1/nodes/N1/suspend" {
			t.Errorf("URL is invalid: %v", req.URL)
		}
		if req.Method != "POST" {
			t.Errorf("Method is invalid: %v", req.Method)
		}

		return c.CreateResponse(200, nbody), nil
	}

	n, err := SuspendNode(c, "P1", "N1")
	if err != nil {
		t.Errorf("Error is invalid: %v", err)
	}

	if n.Name != nbody.Name {
		t.Errorf("Name is invalid: %v", n.Name)
	}

	if count != 1 {
		t.Errorf("Not enough requests: %v", count)
	}
}
