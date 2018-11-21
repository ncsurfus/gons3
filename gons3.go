package gons3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// GetProjects retrieves a list of the projects on the GNS3 server.
func (conn Gns3Conn) GetProjects() (projects []Project, err error) {
	r, err := http.Get(conn.url + "/v2/projects")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get projects")
	}
	if r.StatusCode != 200 {
		return nil, errors.Wrapf(errors.New("http status code "+r.Status), "failed to get projects")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get projects")
	}

	projects = []Project{}
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get projects")
	}

	return projects, nil
}

// GetNodes retrieves a list of nodes from the given project on the GNS server.
func (conn Gns3Conn) GetNodes(projectID string) (nodes []Node, err error) {
	r, err := http.Get(fmt.Sprintf("%s/v2/projects/%s/nodes", conn.url, projectID))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get nodes")
	}
	if r.StatusCode != 200 {
		return nil, errors.Wrapf(errors.New("http status code "+r.Status), "failed to get nodes")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get nodes")
	}

	nodes = []Node{}
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get nodes")
	}

	return nodes, nil
}

// Connect creates a GNS3 connection.
func (conn *Gns3Conn) Connect(url string) (err error) {
	if conn.url != "" {
		return errors.Wrapf(errors.New("already connected to "+conn.url), "failed to connect")
	}

	if url == "" {
		return errors.Wrapf(errors.New("url is empty "+conn.url), "failed to connect")
	}

	r, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "failed to connect")
	}

	if r.StatusCode != 200 {
		return errors.Wrapf(errors.New("http status code "+r.Status), "failed to get nodes")
	}

	conn.url = url
	return nil
}

// URL returns a string of the connected server or an empty string if it is not connected.
func (conn Gns3Conn) URL() (url string) {
	return conn.url
}

// Gns3Conn represents a connection to a GNS3 server.
type Gns3Conn struct {
	url string
}

// Project models a GNS3 Project
type Project struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
}

// Node models a GNS3 Node
type Node struct {
	ComputeID   string `json:"compute_id"`
	ConsolePort int    `json:"console"`
	ConsoleHost string `json:"console_host"`
	ConsoleType string `json:"console_type"`
	Name        string `json:"name"`
	ID          string `json:"node_id"`
	Type        string `json:"node_type"`
	ProjectID   string `json:"project_id"`
	Status      string `json:"status"`
}
