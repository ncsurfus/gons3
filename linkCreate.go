package gons3

// LinkNodeCreate creates a new LinkNode.
type LinkNodeCreate struct {
	NodeID        string
	AdapterNumber int
	PortNumber    int
	label         map[string]interface{}
}

// SetLabelProperty sets a custom property and value for the node.label.
func (n *LinkNodeCreate) SetLabelProperty(name string, value interface{}) {
	if n.label == nil {
		n.label = map[string]interface{}{}
	}
	n.label[name] = value
}

// SetLabelText sets the text for the new linkNode's label.
func (n *LinkNodeCreate) SetLabelText(text string) {
	n.SetLabelProperty("text", text)
}

// SetLabelStyle sets the style for the new linkNode's label.
func (n *LinkNodeCreate) SetLabelStyle(style string) {
	n.SetLabelProperty("style", style)
}

// SetLabelX sets the x position for the new linkNode's label.
func (n *LinkNodeCreate) SetLabelX(x int) {
	n.SetLabelProperty("x", x)
}

// SetLabelY sets the y position for the new linkNode's label.
func (n *LinkNodeCreate) SetLabelY(y int) {
	n.SetLabelProperty("y", y)
}

// SetLabelRotation sets the rotation for the new node's label.
func (n *LinkNodeCreate) SetLabelRotation(rotation int) {
	n.SetLabelProperty("rotation", rotation)
}

// LinkCreate models a new GNS3 link between two or more nodes.
// GNS3 schema requires values: Name, NodeType, and ComputeID
type LinkCreate struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *LinkCreate) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetLinkType sets the type of link.
// Values: ethernet or serial
func (n *LinkCreate) SetLinkType(linkType string) {
	n.SetProperty("link_type", linkType)
}

// SetText set's the links text.
// Nodes text should match the node name!
func (n *LinkCreate) SetText(text string) {
	n.SetProperty("text", text)
}

// SetSuspend sets the suspended status of the link.
func (n *LinkCreate) SetSuspend(isSuspended bool) {
	n.SetProperty("suspend", isSuspended)
}

// SetNodes sets the nodes that are a part of the link.
func (n *LinkCreate) SetNodes(linkNodes []LinkNodeCreate) {
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
