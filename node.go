// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/schemas/node.py
// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/handlers/api/controller/node_handler.py

package gons3

import (
	"errors"
	"net/url"
)

// ErrEmptyNodeID means that the node id cannot be empty.
var ErrEmptyNodeID = errors.New("nodeID cannot be empty")

// Node models a GNS3 node.
type Node struct {
	ComputeID        string                 `json:"compute_id"`
	ProjectID        string                 `json:"project_id"`
	NodeID           string                 `json:"node_id"`
	TemplateID       string                 `json:"template_id"`
	NodeType         string                 `json:"node_type"`
	NodeDirectory    string                 `json:"node_directory"`
	CommandLine      string                 `json:"command_line"`
	Name             string                 `json:"name"`
	Console          int                    `json:"console"`
	ConsoleHost      string                 `json:"console_host"`
	ConsoleType      string                 `json:"console_type"`
	ConsoleAutoStart bool                   `json:"console_auto_start"`
	Properties       map[string]interface{} `json:"properties"`
	Status           string                 `json:"status"`
	Label            Label                  `json:"label"`
	Symbol           string                 `json:"symbol"`
	Width            int                    `json:"width"`
	Height           int                    `json:"height"`
	X                int                    `json:"x"`
	Y                int                    `json:"y"`
	Z                int                    `json:"z"`
	Locked           bool                   `json:"locked"`
	PortNameFormat   string                 `json:"port_name_format"`
	PortSegmentSize  int                    `json:"port_segment_size"`
	FirstPortName    string                 `json:"first_port_name"`
	CustomAdapters   []CustomAdapter        `json:"custom_adapters"`
	Ports            []NodePort             `json:"ports"`
}

// NodePort models a GNS3 node's port.
type NodePort struct {
	Name          string                 `json:"name"`
	ShortName     string                 `json:"short_name"`
	AdapterNumber int                    `json:"adapter_number"`
	AdapterType   string                 `json:"adapter_type"`
	PortNumber    int                    `json:"port_number"`
	LinkType      string                 `json:"link_type"`
	DataLinkTypes map[string]interface{} `json:"data_link_types"`
	MACAddress    string                 `json:"mac_address"`
}

// IsStarted returns true if the node is started, or false if not started.
func (node Node) IsStarted() bool {
	return node.Status == "started"
}

// IsSuspended returns true if the node is suspended, or false if not suspended.
func (node Node) IsSuspended() bool {
	return node.Status == "suspended"
}

// IsStopped returns true if the node is stopped, or false not stopped.
func (node Node) IsStopped() bool {
	return node.Status == "stopped"
}

// CreateNode creates a GNS3 node.
func CreateNode(client GNS3Client, projectID string, nodeBuilder NodeBuilder) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes"
	node := Node{}
	if err := post(client, path, 201, nodeBuilder.values, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// UpdateNode updates a GNS3 node.
func UpdateNode(client GNS3Client, projectID, nodeID string, nodeBuilder NodeUpdater) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID)
	node := Node{}
	if err := put(client, path, 200, nodeBuilder.values, &node); err != nil {
		return Node{}, err
	}

	return node, nil
}

// GetNodes gets all the nodes in the specified project
func GetNodes(client GNS3Client, projectID string) ([]Node, error) {
	if projectID == "" {
		return []Node{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes"
	node := []Node{}
	if err := get(client, path, 200, &node); err != nil {
		return []Node{}, err
	}
	return node, nil
}

// GetNode gets a GNS3 node instance with the specified id.
func GetNode(client GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID)
	node := Node{}
	if err := get(client, path, 200, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// DeleteNode deletes a GNS3 node instance.
func DeleteNode(client GNS3Client, projectID, nodeID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}
	if nodeID == "" {
		return ErrEmptyNodeID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID)
	if err := delete(client, path, 204, nil); err != nil {
		return err
	}
	return nil
}

// StartNode starts a node in a GNS3 project
func StartNode(client GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/start"
	if err := post(client, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// StartNodes starts all nodes in a GNS3 project
func StartNodes(client GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/start"
	if err := post(client, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// StopNode stops a node in a GNS3 project
func StopNode(client GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/stop"
	if err := post(client, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// StopNodes stops all nodes in a GNS3 project
func StopNodes(client GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/stop"
	if err := post(client, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// SuspendNode suspends a node in a GNS3 project
func SuspendNode(client GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/suspend"
	if err := post(client, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// SuspendNodes suspends all nodes in a GNS3 project
func SuspendNodes(client GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/suspend"
	if err := post(client, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// ReloadNode reloads a node in a GNS3 project
func ReloadNode(client GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/reload"
	if err := post(client, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// ReloadNodes restarts all nodes in a GNS3 project
func ReloadNodes(client GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/reload"
	if err := post(client, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// ReadNodeFile reads a GNS3 node's file.
func ReadNodeFile(client GNS3Client, projectID, nodeID, filepath string) ([]byte, error) {
	if projectID == "" {
		return []byte{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return []byte{}, ErrEmptyNodeID
	}
	if filepath == "" {
		return []byte{}, ErrEmptyFilepath
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) +
		"/files/" + filepath
	data := []byte{}
	if err := get(client, path, 200, &data); err != nil {
		return []byte{}, err
	}
	return data, nil
}

// WriteNodeFile writes a GNS3 node's file.
func WriteNodeFile(client GNS3Client, projectID, nodeID, filepath string, data []byte) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}
	if nodeID == "" {
		return ErrEmptyNodeID
	}
	if filepath == "" {
		return ErrEmptyFilepath
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) +
		"/files/" + filepath
	if err := post(client, path, 201, &data, nil); err != nil {
		return err
	}
	return nil
}
