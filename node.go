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

// IsStarted returns true if the node is started, or false if the node is stopped/suspended.
func (n Node) IsStarted() bool {
	return n.Status == "started"
}

// IsSuspended returns true if the node is suspended, or false if the node is started/stopped.
func (n Node) IsSuspended() bool {
	return n.Status == "suspended"
}

// IsStopped returns true if the node is stopped, or false if the node is started/suspended.
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
func UpdateNode(g GNS3Client, projectID string, nodeID string, n NodeUpdater) (Node, error) {
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

// StartNode starts a node in a GNS3 project
func StartNode(g GNS3Client, projectID string, nodeID string) (Node, error) {
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
func StopNode(g GNS3Client, projectID string, nodeID string) (Node, error) {
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
func SuspendNode(g GNS3Client, projectID string, nodeID string) (Node, error) {
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
func ReloadNode(g GNS3Client, projectID string, nodeID string) (Node, error) {
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

// SetConsole sets the console port to the node.
func (n *NodeCreator) SetConsole(port int) {
	n.SetProperty("console", port)
}

// SetConsoleType sets the type of console to use the node.
// vnc, telnet, http, https, spice, spice+agent, none
// ConsoleType is not applicable to the following node types and will be ignored.
//   cloud, nat, ethernet_hub, frame_relay_switch, atm_switch
//   https://github.com/GNS3/gns3-server/blob/2.2/gns3server/controller/node.py#L476
func (n *NodeCreator) SetConsoleType(consoleType string) {
	n.SetProperty("console_type", consoleType)
}

// SetConsoleAutoStart sets if the console should automatically be started.
func (n *NodeCreator) SetConsoleAutoStart(autoStart bool) {
	n.SetProperty("console_auto_start", autoStart)
}

// SetSymbol sets the symbol of the node.
func (n *NodeCreator) SetSymbol(symbol string) {
	n.SetProperty("symbol", symbol)
}

// SetLabel sets the label of the node.
func (n *NodeCreator) SetLabel(label LabelCreator) {
	n.SetProperty("label", label.values)
}

// SetX sets the X position of the node.
func (n *NodeCreator) SetX(x int) {
	n.SetProperty("x", x)
}

// SetY sets the Y position of the node.
func (n *NodeCreator) SetY(y int) {
	n.SetProperty("y", y)
}

// SetZ sets the Z position of the node.
func (n *NodeCreator) SetZ(z int) {
	n.SetProperty("z", z)
}

// SetLocked sets if the node is locked or not.
func (n *NodeCreator) SetLocked(locked bool) {
	n.SetProperty("locked", locked)
}

// SetPortNameFormat sets if the format for the port name. {0} gets replaced with the port number.
func (n *NodeCreator) SetPortNameFormat(format string) {
	n.SetProperty("port_name_format", format)
}

// SetPortSegmentSize sets if the size of the port segment.
func (n *NodeCreator) SetPortSegmentSize(size int) {
	n.SetProperty("port_segment_size", size)
}

// SetFirstPortName sets the name of the first port.
func (n *NodeCreator) SetFirstPortName(name string) {
	n.SetProperty("first_port_name", name)
}

// NodeUpdater models a GNS3 node update.
type NodeUpdater struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *NodeUpdater) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetName sets the name for the new node.
func (n *NodeUpdater) SetName(name string) {
	n.SetProperty("name", name)
}

// SetConsole sets the console port to the node.
func (n *NodeUpdater) SetConsole(port int) {
	n.SetProperty("console", port)
}

// SetConsoleType sets the type of console to use the node.
// vnc, telnet, http, https, spice, spice+agent, none
// ConsoleType is not applicable to the following node types and will be ignored.
//   cloud, nat, ethernet_hub, frame_relay_switch, atm_switch
//   https://github.com/GNS3/gns3-server/blob/2.2/gns3server/controller/node.py#L476
func (n *NodeUpdater) SetConsoleType(consoleType string) {
	n.SetProperty("console_type", consoleType)
}

// SetConsoleAutoStart sets if the console should automatically be started.
func (n *NodeUpdater) SetConsoleAutoStart(autoStart bool) {
	n.SetProperty("console_auto_start", autoStart)
}

// SetSymbol sets the symbol of the node.
func (n *NodeUpdater) SetSymbol(symbol string) {
	n.SetProperty("symbol", symbol)
}

// SetLabel sets the label of the node.
func (n *NodeUpdater) SetLabel(label LabelCreator) {
	n.SetProperty("label", label.values)
}

// SetX sets the X position of the node.
func (n *NodeUpdater) SetX(x int) {
	n.SetProperty("x", x)
}

// SetY sets the Y position of the node.
func (n *NodeUpdater) SetY(y int) {
	n.SetProperty("y", y)
}

// SetZ sets the Z position of the node.
func (n *NodeUpdater) SetZ(z int) {
	n.SetProperty("z", z)
}

// SetLocked sets if the node is locked or not.
func (n *NodeUpdater) SetLocked(locked bool) {
	n.SetProperty("locked", locked)
}

// SetPortNameFormat sets if the format for the port name. {0} gets replaced with the port number.
func (n *NodeUpdater) SetPortNameFormat(format string) {
	n.SetProperty("port_name_format", format)
}

// SetPortSegmentSize sets if the size of the port segment.
func (n *NodeUpdater) SetPortSegmentSize(size int) {
	n.SetProperty("port_segment_size", size)
}

// SetFirstPortName sets the name of the first port.
func (n *NodeUpdater) SetFirstPortName(name string) {
	n.SetProperty("first_port_name", name)
}
