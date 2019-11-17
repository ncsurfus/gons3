package gons3


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
