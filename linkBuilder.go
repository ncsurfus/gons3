package gons3

// LinkNodeBuilder creates a new LinkNode.
type LinkNodeBuilder struct {
	NodeID        string
	AdapterNumber int
	PortNumber    int
	label         map[string]interface{}
}

// SetLabelProperty sets a custom property and value for the node.label.
func (n *LinkNodeBuilder) SetLabelProperty(name string, value interface{}) {
	if n.label == nil {
		n.label = map[string]interface{}{}
	}
	n.label[name] = value
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
func (n *LinkBuilder) SetNodes(linkNodes []LinkNodeBuilder) {
	nodes := make([]map[string]interface{}, len(linkNodes))
	for i, linkNode := range linkNodes {
		node := map[string]interface{}{}
		node["node_id"] = linkNode.NodeID
		node["adapter_number"] = linkNode.AdapterNumber
		node["port_number"] = linkNode.PortNumber
		if linkNode.label != nil {
			node["label"] = linkNode.label
		}
		nodes[i] = node
	}
	n.SetProperty("nodes", nodes)
}
