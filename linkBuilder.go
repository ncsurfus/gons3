package gons3

// LinkNodeBuilder creates a new LinkNode.
type LinkNodeBuilder struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *LinkNodeBuilder) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetLabelProperty sets a custom property and value for the node's label.
func (n *LinkNodeBuilder) SetLabelProperty(name string, value interface{}) {
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

// SetNodeID sets the NodeID for the new linkNode.
func (n *LinkNodeBuilder) SetNodeID(nodeID string) {
	n.SetProperty("node_id", nodeID)
}

// SetAdapterNumber sets the adapter number for the new linkNode.
func (n *LinkNodeBuilder) SetAdapterNumber(adapterNumber int) {
	n.SetProperty("adapter_number", adapterNumber)
}

// SetPortNumber sets the port number for the new linkNode.
func (n *LinkNodeBuilder) SetPortNumber(portNumber int) {
	n.SetProperty("port_number", portNumber)
}

// SetLabelText sets the text for the new linkNode's label.
func (n *LinkNodeBuilder) SetLabelText(text string) {
	n.SetLabelProperty("text", text)
}

// SetLabelStyle sets the style for the new linkNode's label.
func (n *LinkNodeBuilder) SetLabelStyle(style string) {
	n.SetLabelProperty("style", style)
}

// SetLabelX sets the x position for the new linkNode's label.
func (n *LinkNodeBuilder) SetLabelX(x int) {
	n.SetLabelProperty("x", x)
}

// SetLabelY sets the y position for the new linkNode's label.
func (n *LinkNodeBuilder) SetLabelY(y int) {
	n.SetLabelProperty("y", y)
}

// SetLabelRotation sets the rotation for the new node's label.
func (n *LinkNodeBuilder) SetLabelRotation(rotation int) {
	n.SetLabelProperty("rotation", rotation)
}

// LinkBuilder models a new GNS3 link between two or more nodes.
// GNS3 schema requires values: Name, NodeType, and ComputeID
type LinkBuilder struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *LinkBuilder) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetLinkType sets the type of link.
// Values: ethernet or serial
func (n *LinkBuilder) SetLinkType(linkType string) {
	n.SetProperty("link_type", linkType)
}

// SetText set's the links text.
// Nodes text should match the node name!
func (n *LinkBuilder) SetText(text string) {
	n.SetProperty("text", text)
}

// SetSuspend sets the suspended status of the link.
func (n *LinkBuilder) SetSuspend(isSuspended bool) {
	n.SetProperty("suspend", isSuspended)
}

// SetNodes sets the nodes that are a part of the link.
func (n *LinkBuilder) SetNodes(linkNodeBuilders []LinkNodeBuilder) {
	linkNodes := make([]map[string]interface{}, len(linkNodeBuilders))
	for i, linkNode := range linkNodeBuilders {
		linkNodes[i] = linkNode.values
	}
	n.SetProperty("nodes", linkNodes)
}
