package gons3

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
