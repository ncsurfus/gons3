package gons3

// LinkNodeUpdater creates a new LinkNode.
type LinkNodeUpdater struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *LinkNodeUpdater) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetLabelProperty sets a custom property and value for the node's label.
func (n *LinkNodeUpdater) SetLabelProperty(name string, value interface{}) {
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

// SetNodeID sets the NodeID for the node to be updated.
func (n *LinkNodeUpdater) SetNodeID(nodeID string) {
	n.SetProperty("adapter_number", 0)
	n.SetProperty("port_number", 0)
	n.SetProperty("node_id", nodeID)
}

// SetLabelText sets the updated text for the link node.
func (n *LinkNodeUpdater) SetLabelText(text string) {
	n.SetLabelProperty("text", text)
}

// SetLabelStyle sets the updated style for the link node.
func (n *LinkNodeUpdater) SetLabelStyle(style string) {
	n.SetLabelProperty("style", style)
}

// SetLabelX sets the updated x position for the link node.
func (n *LinkNodeUpdater) SetLabelX(x int) {
	n.SetLabelProperty("x", x)
}

// SetLabelY sets the updated y position for the link node.
func (n *LinkNodeUpdater) SetLabelY(y int) {
	n.SetLabelProperty("y", y)
}

// SetLabelRotation sets the updated rotation for the link node.
func (n *LinkNodeUpdater) SetLabelRotation(rotation int) {
	n.SetLabelProperty("rotation", rotation)
}

// LinkUpdater models an update to a GNS3 link between two or more nodes.
type LinkUpdater struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *LinkUpdater) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetSuspend sets the suspended status of the link.
func (n *LinkUpdater) SetSuspend(isSuspended bool) {
	n.SetProperty("suspend", isSuspended)
}

// SetNodes sets the nodes that are a part of the link.
func (n *LinkUpdater) SetNodes(linkNodeUpdater ...LinkNodeUpdater) {
	linkNodes := make([]map[string]interface{}, len(linkNodeUpdater))
	for i, linkNode := range linkNodeUpdater {
		linkNodes[i] = linkNode.values
	}
	n.SetProperty("nodes", linkNodes)
}
