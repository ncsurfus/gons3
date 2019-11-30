package gns3tests

import (
	"gons3"
	"testing"
)

func TestGetLink(t *testing.T) {
	ProjectBuilder := gons3.ProjectBuilder{}
	ProjectBuilder.SetName("TestGetLink")
	project, err := gons3.CreateProject(client, ProjectBuilder)
	if err != nil {
		fatalAssert(t, "CreateProject error", nil, err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NodeBuilder{}
	nodeBuilderA.SetName("TheNodeA")
	nodeBuilderA.SetNodeType("vpcs")
	nodeBuilderA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		fatalAssert(t, "CreateNode(A) error", nil, err)
	}

	nodeBuilderB := gons3.NodeBuilder{}
	nodeBuilderB.SetName("TheNodeB")
	nodeBuilderB.SetNodeType("vpcs")
	nodeBuilderB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		fatalAssert(t, "CreateNode(B) error", nil, err)
	}

	LinkNodeBuilderA := gons3.LinkNodeBuilder{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	LinkNodeBuilderB := gons3.LinkNodeBuilder{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkBuilder := gons3.LinkBuilder{}
	linkBuilder.SetNodes([]gons3.LinkNodeBuilder{LinkNodeBuilderA, LinkNodeBuilderB})
	createdLink, err := gons3.CreateLink(client, project.ProjectID, linkBuilder)
	if err != nil {
		fatalAssert(t, "CreateLink error", nil, err)
	}

	link, err := gons3.GetLink(client, project.ProjectID, createdLink.LinkID)
	if err != nil {
		fatalAssert(t, "GetLink error", nil, err)
	}
	fatalAssert(t, "len(link.Nodes)", 2, len(link.Nodes))
	errorAssert(t, "link.Nodes[0].NodeID", nodeA.NodeID, link.Nodes[0].NodeID)
	errorAssert(t, "link.Nodes[0].AdapterNumber", nodeA.Ports[0].AdapterNumber, link.Nodes[0].AdapterNumber)
	errorAssert(t, "link.Nodes[0].PortNumber", nodeA.Ports[0].PortNumber, link.Nodes[0].PortNumber)
	errorAssert(t, "link.Nodes[1].NodeID", nodeB.NodeID, link.Nodes[1].NodeID)
	errorAssert(t, "link.Nodes[1].AdapterNumber", nodeB.Ports[0].AdapterNumber, link.Nodes[1].AdapterNumber)
	errorAssert(t, "link.Nodes[1].PortNumber", nodeB.Ports[0].PortNumber, link.Nodes[1].PortNumber)
}

func TestGetLinks(t *testing.T) {
	ProjectBuilder := gons3.ProjectBuilder{}
	ProjectBuilder.SetName("TestGetLinks")
	project, err := gons3.CreateProject(client, ProjectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NodeBuilder{}
	nodeBuilderA.SetName("TheNodeA")
	nodeBuilderA.SetNodeType("ethernet_switch")
	nodeBuilderA.SetConsoleType("none")
	nodeBuilderA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeBuilderB := gons3.NodeBuilder{}
	nodeBuilderB.SetName("TheNodeB")
	nodeBuilderB.SetNodeType("ethernet_switch")
	nodeBuilderB.SetConsoleType("none")
	nodeBuilderB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	LinkNodeBuilderA := gons3.LinkNodeBuilder{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	LinkNodeBuilderB := gons3.LinkNodeBuilder{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkBuilderA := gons3.LinkBuilder{}
	linkBuilderA.SetNodes([]gons3.LinkNodeBuilder{LinkNodeBuilderA, LinkNodeBuilderB})
	createdLinkA, err := gons3.CreateLink(client, project.ProjectID, linkBuilderA)
	if err != nil {
		t.Fatalf("Error creating linkA: %v", err)
	}

	LinkNodeBuilderC := gons3.LinkNodeBuilder{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[1].PortNumber,
	}
	LinkNodeBuilderD := gons3.LinkNodeBuilder{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[1].PortNumber,
	}
	linkBuilderB := gons3.LinkBuilder{}
	linkBuilderB.SetNodes([]gons3.LinkNodeBuilder{LinkNodeBuilderC, LinkNodeBuilderD})
	createdLinkB, err := gons3.CreateLink(client, project.ProjectID, linkBuilderB)
	if err != nil {
		t.Fatalf("Error creating linkB: %v", err)
	}

	link, err := gons3.GetLinks(client, project.ProjectID)
	if err != nil {
		t.Fatalf("Error getting link: %v", err)
	}
	if len(link) != 2 {
		t.Fatalf("Invalid link count: %v", len(link))
	}
	if link[0].LinkID != createdLinkA.LinkID {
		t.Errorf("Invalid link.0.linkID: %v", createdLinkA.LinkID)
	}
	if link[1].LinkID != createdLinkB.LinkID {
		t.Errorf("Invalid link.1.linkID: %v", createdLinkB.LinkID)
	}
}

func TestGetNodeLinks(t *testing.T) {
	ProjectBuilder := gons3.ProjectBuilder{}
	ProjectBuilder.SetName("TestGetNodeLinks")
	project, err := gons3.CreateProject(client, ProjectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NodeBuilder{}
	nodeBuilderA.SetName("TheNodeA")
	nodeBuilderA.SetNodeType("ethernet_switch")
	nodeBuilderA.SetConsoleType("none")
	nodeBuilderA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeBuilderB := gons3.NodeBuilder{}
	nodeBuilderB.SetName("TheNodeB")
	nodeBuilderB.SetNodeType("ethernet_switch")
	nodeBuilderB.SetConsoleType("none")
	nodeBuilderB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	LinkNodeBuilderA := gons3.LinkNodeBuilder{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	LinkNodeBuilderB := gons3.LinkNodeBuilder{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkBuilderA := gons3.LinkBuilder{}
	linkBuilderA.SetNodes([]gons3.LinkNodeBuilder{LinkNodeBuilderA, LinkNodeBuilderB})
	createdLinkA, err := gons3.CreateLink(client, project.ProjectID, linkBuilderA)
	if err != nil {
		t.Fatalf("Error creating linkA: %v", err)
	}

	LinkNodeBuilderC := gons3.LinkNodeBuilder{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[1].PortNumber,
	}
	LinkNodeBuilderD := gons3.LinkNodeBuilder{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[1].PortNumber,
	}
	linkBuilderB := gons3.LinkBuilder{}
	linkBuilderB.SetNodes([]gons3.LinkNodeBuilder{LinkNodeBuilderC, LinkNodeBuilderD})
	createdLinkB, err := gons3.CreateLink(client, project.ProjectID, linkBuilderB)
	if err != nil {
		t.Fatalf("Error creating linkB: %v", err)
	}

	link, err := gons3.GetNodeLinks(client, project.ProjectID, nodeA.NodeID)
	if err != nil {
		t.Fatalf("Error getting node links: %v", err)
	}
	if len(link) != 2 {
		t.Fatalf("Invalid link count: %v", len(link))
	}
	if link[0].LinkID != createdLinkA.LinkID {
		t.Errorf("Invalid link.0.linkID: %v", createdLinkA.LinkID)
	}
	if link[1].LinkID != createdLinkB.LinkID {
		t.Errorf("Invalid link.1.linkID: %v", createdLinkB.LinkID)
	}
}

func TestDeleteLink(t *testing.T) {
	ProjectBuilder := gons3.ProjectBuilder{}
	ProjectBuilder.SetName("TestDeleteLink")
	project, err := gons3.CreateProject(client, ProjectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NodeBuilder{}
	nodeBuilderA.SetName("TheNodeA")
	nodeBuilderA.SetNodeType("vpcs")
	nodeBuilderA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeBuilderB := gons3.NodeBuilder{}
	nodeBuilderB.SetName("TheNodeB")
	nodeBuilderB.SetNodeType("vpcs")
	nodeBuilderB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	LinkNodeBuilderA := gons3.LinkNodeBuilder{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	LinkNodeBuilderB := gons3.LinkNodeBuilder{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkBuilder := gons3.LinkBuilder{}
	linkBuilder.SetNodes([]gons3.LinkNodeBuilder{LinkNodeBuilderA, LinkNodeBuilderB})
	createdLink, err := gons3.CreateLink(client, project.ProjectID, linkBuilder)
	if err != nil {
		t.Fatalf("Error creating link: %v", err)
	}

	err = gons3.DeleteLink(client, project.ProjectID, createdLink.LinkID)
	if err != nil {
		t.Fatalf("Error deleting link: %v", err)
	}
}
