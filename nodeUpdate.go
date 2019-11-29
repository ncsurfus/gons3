package gons3

// NodeUpdate models a GNS3 node update.
type NodeUpdate struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *NodeUpdate) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetName sets the name for the new node.
func (n *NodeUpdate) SetName(name string) {
	n.SetProperty("name", name)
}

// SetConsole sets the console port to the node.
func (n *NodeUpdate) SetConsole(port int) {
	n.SetProperty("console", port)
}

// SetConsoleType sets the type of console to use the node.
// vnc, telnet, http, https, spice, spice+agent, none
// ConsoleType is not applicable to the following node types and will be ignored.
//   cloud, nat, ethernet_hub, frame_relay_switch, atm_switch
//   https://github.com/GNS3/gns3-server/blob/2.2/gns3server/controller/node.py#L476
func (n *NodeUpdate) SetConsoleType(consoleType string) {
	n.SetProperty("console_type", consoleType)
}

// SetConsoleAutoStart sets if the console should automatically be started.
func (n *NodeUpdate) SetConsoleAutoStart(autoStart bool) {
	n.SetProperty("console_auto_start", autoStart)
}

// SetSymbol sets the symbol of the node.
func (n *NodeUpdate) SetSymbol(symbol string) {
	n.SetProperty("symbol", symbol)
}

// SetX sets the X position of the node.
func (n *NodeUpdate) SetX(x int) {
	n.SetProperty("x", x)
}

// SetY sets the Y position of the node.
func (n *NodeUpdate) SetY(y int) {
	n.SetProperty("y", y)
}

// SetZ sets the Z position of the node.
func (n *NodeUpdate) SetZ(z int) {
	n.SetProperty("z", z)
}

// SetLocked sets if the node is locked or not.
func (n *NodeUpdate) SetLocked(locked bool) {
	n.SetProperty("locked", locked)
}

// SetPortNameFormat sets if the format for the port name. {0} gets replaced with the port number.
func (n *NodeUpdate) SetPortNameFormat(format string) {
	n.SetProperty("port_name_format", format)
}

// SetPortSegmentSize sets if the size of the port segment.
func (n *NodeUpdate) SetPortSegmentSize(size int) {
	n.SetProperty("port_segment_size", size)
}

// SetFirstPortName sets the name of the first port.
func (n *NodeUpdate) SetFirstPortName(name string) {
	n.SetProperty("first_port_name", name)
}

// SetLabelProperty sets a custom property and value for the node.label.
func (n *NodeUpdate) SetLabelProperty(name string, value interface{}) {
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
func (n *NodeUpdate) SetLabelStyle(style string) {
	n.SetLabelProperty("style", style)
}

// SetLabelX sets the x position for the new node's label.
func (n *NodeUpdate) SetLabelX(x int) {
	n.SetLabelProperty("x", x)
}

// SetLabelY sets the y position for the new node's label.
func (n *NodeUpdate) SetLabelY(y int) {
	n.SetLabelProperty("y", y)
}

// SetLabelRotation sets the rotation for the new node's label.
func (n *NodeUpdate) SetLabelRotation(rotation int) {
	n.SetLabelProperty("rotation", rotation)
}
