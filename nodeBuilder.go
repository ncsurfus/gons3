package gons3

// IsConsoleSupport checks if the type of node supports console.
func IsConsoleSupport(nodeType string) bool {
	switch nodeType {
	case "cloud", "nat", "ethernet_switch", "ethernet_hub", "frame_relay_switch", "atm_switch":
		return false
	default:
		return true
	}
}

// NodeBuilder models a new GNS3 node.
// GNS3 schema requires values: Name, NodeType, and ComputeID
type NodeBuilder struct {
	values map[string]interface{}
}

// NewNodeBuilder initializes a new NodeBuilder with the required values and local compute.
func NewNodeBuilder(name, nodeType string) NodeBuilder {
	nodeBuilder := NodeBuilder{}
	nodeBuilder.SetLocalComputeID()
	nodeBuilder.SetName(name)
	nodeBuilder.SetNodeType(nodeType)

	if !IsConsoleSupport(nodeType) {
		nodeBuilder.SetConsoleType("none")
	}
	return nodeBuilder
}

// SetProperty sets a custom property and value for the node.
func (n *NodeBuilder) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetName sets the name for the new node.
func (n *NodeBuilder) SetName(name string) {
	n.SetProperty("name", name)
}

// SetNodeType sets the node type for the new node.
// GNS3 schema provides these values: cloud, nat, ethernet_hub, ethernet_switch,
// frame_relay_switch, atm_switch, docker, dynamips, vpcs, traceng, virtualbox, vmware, iou, qemu
func (n *NodeBuilder) SetNodeType(nodeType string) {
	n.SetProperty("node_type", nodeType)
}

// SetComputeID sets the compute_id for the new node.
// See SetLocalComputeID() to set the the compute_id to the local instance.
func (n *NodeBuilder) SetComputeID(id string) {
	n.SetProperty("compute_id", id)
}

// SetLocalComputeID sets the compute_id to local for the new node.
func (n *NodeBuilder) SetLocalComputeID() {
	n.SetProperty("compute_id", "local")
}

// SetConsole sets the console port to the node.
func (n *NodeBuilder) SetConsole(port int) {
	n.SetProperty("console", port)
}

// SetConsoleType sets the type of console to use the node.
// vnc, telnet, http, https, spice, spice+agent, none
// ConsoleType is not applicable to the following node types and will be ignored.
//   cloud, nat, ethernet_hub, frame_relay_switch, atm_switch
//   https://github.com/GNS3/gns3-server/blob/2.2/gns3server/controller/node.py#L476
func (n *NodeBuilder) SetConsoleType(consoleType string) {
	n.SetProperty("console_type", consoleType)
}

// SetConsoleAutoStart sets if the console should automatically be started.
func (n *NodeBuilder) SetConsoleAutoStart(autoStart bool) {
	n.SetProperty("console_auto_start", autoStart)
}

// SetSymbol sets the symbol of the node.
func (n *NodeBuilder) SetSymbol(symbol string) {
	n.SetProperty("symbol", symbol)
}

// SetX sets the X position of the node.
func (n *NodeBuilder) SetX(x int) {
	n.SetProperty("x", x)
}

// SetY sets the Y position of the node.
func (n *NodeBuilder) SetY(y int) {
	n.SetProperty("y", y)
}

// SetZ sets the Z position of the node.
func (n *NodeBuilder) SetZ(z int) {
	n.SetProperty("z", z)
}

// SetLocked sets if the node is locked or not.
func (n *NodeBuilder) SetLocked(locked bool) {
	n.SetProperty("locked", locked)
}

// SetPortNameFormat sets if the format for the port name. {0} gets replaced with the port number.
func (n *NodeBuilder) SetPortNameFormat(format string) {
	n.SetProperty("port_name_format", format)
}

// SetPortSegmentSize sets if the size of the port segment.
func (n *NodeBuilder) SetPortSegmentSize(size int) {
	n.SetProperty("port_segment_size", size)
}

// SetFirstPortName sets the name of the first port.
func (n *NodeBuilder) SetFirstPortName(name string) {
	n.SetProperty("first_port_name", name)
}

// SetLabelProperty sets a custom property and value for the node's label.
func (n *NodeBuilder) SetLabelProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	label, ok := n.values["label"].(map[string]interface{})
	if !ok {
		label = map[string]interface{}{
			"text": "",
		}
		n.values["label"] = label
	}
	label[name] = value
}

// SetLabelStyle sets the style for the new node's label.
func (n *NodeBuilder) SetLabelStyle(style string) {
	n.SetLabelProperty("style", style)
}

// SetLabelX sets the x position for the new node's label.
func (n *NodeBuilder) SetLabelX(x int) {
	n.SetLabelProperty("x", x)
}

// SetLabelY sets the y position for the new node's label.
func (n *NodeBuilder) SetLabelY(y int) {
	n.SetLabelProperty("y", y)
}

// SetLabelRotation sets the rotation for the new node's label.
func (n *NodeBuilder) SetLabelRotation(rotation int) {
	n.SetLabelProperty("rotation", rotation)
}
