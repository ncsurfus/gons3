package gons3

// LinkNodeCreator creates a new LinkNode.
type LinkNodeCreator struct {
	NodeID        string
	AdapterNumber int
	PortNumber    int
	Label         LabelCreator
}

// LinkCreator models a new GNS3 link between two or more nodes.
// GNS3 schema requires values: Name, NodeType, and ComputeID
type LinkCreator struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *LinkCreator) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetLinkType sets the type of link.
// Values: ethernet or serial
func (n *LinkCreator) SetLinkType(linkType string) {
	n.SetProperty("link_type", linkType)
}

// SetSuspend sets the suspended status of the link.
func (n *LinkCreator) SetSuspend(isSuspended bool) {
	n.SetProperty("suspend", isSuspended)
}

// SetNodes sets the nodes that are a part of the link.
func (n *LinkCreator) SetNodes(nodeCreators []LinkNodeCreator) {
	nodes := make([]map[string]interface{}, len(nodeCreators))
	for i, node := range nodeCreators {
		n := map[string]interface{}{}
		n["node_id"] = node.NodeID
		n["adapter_number"] = node.AdapterNumber
		n["port_number"] = node.PortNumber
		if node.Label.values != nil {
			n["label"] = node.Label.values
		}
		nodes[i] = n
	}
	n.SetProperty("nodes", nodes)
}
