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
func (n Node) IsStarted() bool {
	return n.Status == "started"
}

// IsSuspended returns true if the node is suspended, or false if not suspended.
func (n Node) IsSuspended() bool {
	return n.Status == "suspended"
}

// IsStopped returns true if the node is stopped, or false not stopped.
func (n Node) IsStopped() bool {
	return n.Status == "stopped"
}

// CreateNode creates a GNS3 node.
func CreateNode(g GNS3Client, projectID string, n NodeCreator) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes"
	node := Node{}
	if err := post(g, path, 201, n.values, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// UpdateNode updates a GNS3 node.
func UpdateNode(g GNS3Client, projectID, nodeID string, n NodeUpdater) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID)
	node := Node{}
	if err := put(g, path, 200, n.values, &node); err != nil {
		return Node{}, err
	}

	return node, nil
}

// GetNodes gets all the nodes in the specified project
func GetNodes(g GNS3Client, projectID string) ([]Node, error) {
	if projectID == "" {
		return []Node{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes"
	node := []Node{}
	if err := get(g, path, 200, &node); err != nil {
		return []Node{}, err
	}
	return node, nil
}

// GetNode gets a GNS3 node instance with the specified id.
func GetNode(g GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID)
	node := Node{}
	if err := get(g, path, 200, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// DeleteNode deletes a GNS3 node instance.
func DeleteNode(g GNS3Client, projectID, nodeID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}
	if nodeID == "" {
		return ErrEmptyNodeID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID)
	if err := delete(g, path, 204, nil); err != nil {
		return err
	}
	return nil
}

// StartNode starts a node in a GNS3 project
func StartNode(g GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/start"
	if err := post(g, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// StartNodes starts all nodes in a GNS3 project
func StartNodes(g GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/start"
	if err := post(g, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// StopNode stops a node in a GNS3 project
func StopNode(g GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/stop"
	if err := post(g, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// StopNodes stops all nodes in a GNS3 project
func StopNodes(g GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/stop"
	if err := post(g, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// SuspendNode suspends a node in a GNS3 project
func SuspendNode(g GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/suspend"
	if err := post(g, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// SuspendNodes suspends all nodes in a GNS3 project
func SuspendNodes(g GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/suspend"
	if err := post(g, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// ReloadNode reloads a node in a GNS3 project
func ReloadNode(g GNS3Client, projectID, nodeID string) (Node, error) {
	if projectID == "" {
		return Node{}, ErrEmptyProjectID
	}
	if nodeID == "" {
		return Node{}, ErrEmptyNodeID
	}

	node := Node{}
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/" + url.PathEscape(nodeID) + "/reload"
	if err := post(g, path, 200, nil, &node); err != nil {
		return Node{}, err
	}
	return node, nil
}

// ReloadNodes restarts all nodes in a GNS3 project
func ReloadNodes(g GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes/reload"
	if err := post(g, path, 204, nil, nil); err != nil {
		return err
	}
	return nil
}

// NodeCreator models a new GNS3 node.
// GNS3 schema requires values: Name, NodeType, and ComputeID
type NodeCreator struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *NodeCreator) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// ReadNodeFile reads a GNS3 node's file.
func ReadNodeFile(g GNS3Client, projectID, nodeID, filepath string) ([]byte, error) {
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
	if err := get(g, path, 200, &data); err != nil {
		return []byte{}, err
	}
	return data, nil
}

// WriteNodeFile writes a GNS3 node's file.
func WriteNodeFile(g GNS3Client, projectID, nodeID, filepath string, data []byte) error {
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
	if err := post(g, path, 201, &data, nil); err != nil {
		return err
	}
	return nil
}
