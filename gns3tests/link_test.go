package gns3tests

import (
	"gons3"
	"testing"
)

func TestGetLink(t *testing.T) {
	ProjectCreate := gons3.ProjectCreate{}
	ProjectCreate.SetName("TestGetLink")
	project, err := gons3.CreateProject(client, ProjectCreate)
	if err != nil {
		fatalAssert(t, "CreateProject error", nil, err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeCreateA := gons3.NodeCreate{}
	nodeCreateA.SetName("TheNodeA")
	nodeCreateA.SetNodeType("vpcs")
	nodeCreateA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeCreateA)
	if err != nil {
		fatalAssert(t, "CreateNode(A) error", nil, err)
	}

	nodeCreateB := gons3.NodeCreate{}
	nodeCreateB.SetName("TheNodeB")
	nodeCreateB.SetNodeType("vpcs")
	nodeCreateB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeCreateB)
	if err != nil {
		fatalAssert(t, "CreateNode(B) error", nil, err)
	}

	linkNodeCreateA := gons3.LinkNodeCreate{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	linkNodeCreateB := gons3.LinkNodeCreate{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkCreate := gons3.LinkCreate{}
	linkCreate.SetNodes([]gons3.LinkNodeCreate{linkNodeCreateA, linkNodeCreateB})
	newLink, err := gons3.CreateLink(client, project.ProjectID, linkCreate)
	if err != nil {
		fatalAssert(t, "CreateLink error", nil, err)
	}

	link, err := gons3.GetLink(client, project.ProjectID, newLink.LinkID)
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
	ProjectCreate := gons3.ProjectCreate{}
	ProjectCreate.SetName("TestGetLinks")
	project, err := gons3.CreateProject(client, ProjectCreate)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeCreateA := gons3.NodeCreate{}
	nodeCreateA.SetName("TheNodeA")
	nodeCreateA.SetNodeType("ethernet_switch")
	nodeCreateA.SetConsoleType("none")
	nodeCreateA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeCreateA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeCreateB := gons3.NodeCreate{}
	nodeCreateB.SetName("TheNodeB")
	nodeCreateB.SetNodeType("ethernet_switch")
	nodeCreateB.SetConsoleType("none")
	nodeCreateB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeCreateB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	linkNodeCreateA := gons3.LinkNodeCreate{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	linkNodeCreateB := gons3.LinkNodeCreate{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkCreateA := gons3.LinkCreate{}
	linkCreateA.SetNodes([]gons3.LinkNodeCreate{linkNodeCreateA, linkNodeCreateB})
	newLinkA, err := gons3.CreateLink(client, project.ProjectID, linkCreateA)
	if err != nil {
		t.Fatalf("Error creating linkA: %v", err)
	}

	linkNodeCreateC := gons3.LinkNodeCreate{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[1].PortNumber,
	}
	linkNodeCreateD := gons3.LinkNodeCreate{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[1].PortNumber,
	}
	linkCreateB := gons3.LinkCreate{}
	linkCreateB.SetNodes([]gons3.LinkNodeCreate{linkNodeCreateC, linkNodeCreateD})
	newLinkB, err := gons3.CreateLink(client, project.ProjectID, linkCreateB)
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
	if link[0].LinkID != newLinkA.LinkID {
		t.Errorf("Invalid link.0.linkID: %v", newLinkA.LinkID)
	}
	if link[1].LinkID != newLinkB.LinkID {
		t.Errorf("Invalid link.1.linkID: %v", newLinkB.LinkID)
	}
}

func TestGetNodeLinks(t *testing.T) {
	ProjectCreate := gons3.ProjectCreate{}
	ProjectCreate.SetName("TestGetNodeLinks")
	project, err := gons3.CreateProject(client, ProjectCreate)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeCreateA := gons3.NodeCreate{}
	nodeCreateA.SetName("TheNodeA")
	nodeCreateA.SetNodeType("ethernet_switch")
	nodeCreateA.SetConsoleType("none")
	nodeCreateA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeCreateA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeCreateB := gons3.NodeCreate{}
	nodeCreateB.SetName("TheNodeB")
	nodeCreateB.SetNodeType("ethernet_switch")
	nodeCreateB.SetConsoleType("none")
	nodeCreateB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeCreateB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	linkNodeCreateA := gons3.LinkNodeCreate{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	linkNodeCreateB := gons3.LinkNodeCreate{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkCreateA := gons3.LinkCreate{}
	linkCreateA.SetNodes([]gons3.LinkNodeCreate{linkNodeCreateA, linkNodeCreateB})
	newLinkA, err := gons3.CreateLink(client, project.ProjectID, linkCreateA)
	if err != nil {
		t.Fatalf("Error creating linkA: %v", err)
	}

	linkNodeCreateC := gons3.LinkNodeCreate{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[1].PortNumber,
	}
	linkNodeCreateD := gons3.LinkNodeCreate{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[1].PortNumber,
	}
	linkCreateB := gons3.LinkCreate{}
	linkCreateB.SetNodes([]gons3.LinkNodeCreate{linkNodeCreateC, linkNodeCreateD})
	newLinkB, err := gons3.CreateLink(client, project.ProjectID, linkCreateB)
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
	if link[0].LinkID != newLinkA.LinkID {
		t.Errorf("Invalid link.0.linkID: %v", newLinkA.LinkID)
	}
	if link[1].LinkID != newLinkB.LinkID {
		t.Errorf("Invalid link.1.linkID: %v", newLinkB.LinkID)
	}
}

func TestDeleteLink(t *testing.T) {
	ProjectCreate := gons3.ProjectCreate{}
	ProjectCreate.SetName("TestDeleteLink")
	project, err := gons3.CreateProject(client, ProjectCreate)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeCreateA := gons3.NodeCreate{}
	nodeCreateA.SetName("TheNodeA")
	nodeCreateA.SetNodeType("vpcs")
	nodeCreateA.SetLocalComputeID()
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeCreateA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeCreateB := gons3.NodeCreate{}
	nodeCreateB.SetName("TheNodeB")
	nodeCreateB.SetNodeType("vpcs")
	nodeCreateB.SetLocalComputeID()
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeCreateB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	linkNodeCreateA := gons3.LinkNodeCreate{
		NodeID:        nodeA.NodeID,
		AdapterNumber: nodeA.Ports[0].AdapterNumber,
		PortNumber:    nodeA.Ports[0].PortNumber,
	}
	linkNodeCreateB := gons3.LinkNodeCreate{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkCreate := gons3.LinkCreate{}
	linkCreate.SetNodes([]gons3.LinkNodeCreate{linkNodeCreateA, linkNodeCreateB})
	newLink, err := gons3.CreateLink(client, project.ProjectID, linkCreate)
	if err != nil {
		t.Fatalf("Error creating link: %v", err)
	}

	err = gons3.DeleteLink(client, project.ProjectID, newLink.LinkID)
	if err != nil {
		t.Fatalf("Error deleting link: %v", err)
	}
}
