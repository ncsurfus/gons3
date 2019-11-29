package gons3

// LinkUpdate models an update to a GNS3 link between two or more nodes.
type LinkUpdate struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the node.
func (n *LinkUpdate) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetSuspend sets the suspended status of the link.
func (n *LinkUpdate) SetSuspend(isSuspended bool) {
	n.SetProperty("suspend", isSuspended)
}

// SetNodes sets the nodes that are a part of the link.
func (n *LinkUpdate) SetNodes(linkNodes []LinkNodeCreate) {
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
