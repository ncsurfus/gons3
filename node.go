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

// CustomAdapter models a GNS3 custom adapter.
type CustomAdapter struct {
	AdapterNumber int    `json:"adapter_number"`
	PortName      string `json:"port_name"`
	AdapterType   string `json:"adapter_type"`
	MACAddress    string `json:"mac_address"`
}

// Label models a GNS3 label.
type Label struct {
	Text     string `json:"text"`
	Style    string `json:"style"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Rotation int    `json:"rotation"`
}

// CreateNode creates a GNS3 node with the specified name.
func CreateNode(g GNS3Client, projectID string, n NodeCreator) (Node, error) {
	path := "/v2/projects/" + url.PathEscape(projectID) + "/nodes"
	node := Node{}
	if err := post(g, path, 201, n.values, &node); err != nil {
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
func GetNode(g GNS3Client, projectID string, nodeID string) (Node, error) {
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
func DeleteNode(g GNS3Client, projectID string, nodeID string) error {
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

// SetName sets the name for the new node.
func (n *NodeCreator) SetName(name string) {
	n.SetProperty("name", name)
}

// SetNodeType sets the node type for the new node.
// GNS3 schema provides these values: cloud, nat, ethernet_hub, ethernet_switch,
// frame_relay_switch, atm_switch, docker, dynamips, vpcs, traceng, virtualbox, vmware, iou, qemu
func (n *NodeCreator) SetNodeType(name string) {
	n.SetProperty("node_type", name)
}

// SetComputeID sets the compute_id for the new node.
// See SetLocalComputeID() to set the the compute_id to the local instance.
func (n *NodeCreator) SetComputeID(id string) {
	n.SetProperty("compute_id", id)
}

// SetLocalComputeID sets the compute_id to local for the new node.
func (n *NodeCreator) SetLocalComputeID() {
	n.SetProperty("compute_id", "local")
}
