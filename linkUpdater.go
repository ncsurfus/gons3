package gons3

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
func (n *LinkUpdater) SetNodes(linkNodeBuilders []LinkNodeBuilder) {
	linkNodes := make([]map[string]interface{}, len(linkNodeBuilders))
	for i, linkNode := range linkNodeBuilders {
		linkNodes[i] = linkNode.values
	}
	n.SetProperty("nodes", linkNodes)
}
