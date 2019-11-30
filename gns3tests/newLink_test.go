package gns3tests

import (
	"gons3"
	"testing"
)

func TestCreateLinkA(t *testing.T) {
	projectBuilder := gons3.ProjectBuilder{}
	projectBuilder.SetName("TestCreateLinkA")
	project, err := gons3.CreateProject(client, projectBuilder)
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
	LinkNodeBuilderA.SetLabelX(5)
	LinkNodeBuilderA.SetLabelY(6)
	LinkNodeBuilderA.SetLabelRotation(90)
	LinkNodeBuilderA.SetLabelStyle("font-family: Verdana;")
	LinkNodeBuilderA.SetLabelText("PortA")

	LinkNodeBuilderB := gons3.LinkNodeBuilder{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	LinkNodeBuilderB.SetLabelX(7)
	LinkNodeBuilderB.SetLabelY(8)
	LinkNodeBuilderB.SetLabelRotation(180)
	LinkNodeBuilderB.SetLabelText("PortB")

	linkBuilder := gons3.LinkBuilder{}
	linkBuilder.SetLinkType("ethernet")
	linkBuilder.SetSuspend(false)
	linkBuilder.SetNodes([]gons3.LinkNodeBuilder{LinkNodeBuilderA, LinkNodeBuilderB})
	link, err := gons3.CreateLink(client, project.ProjectID, linkBuilder)
	if err != nil {
		t.Fatalf("Error creating link: %v", err)
	}

	if link.LinkType != "ethernet" {
		t.Errorf("Invalid link type: %v", link.LinkType)
	}
	if link.Suspended != false {
		t.Errorf("Invalid suspend: %v", link.Suspended)
	}
	if len(link.Nodes) != 2 {
		t.Fatalf("Invalid node count: %v", len(link.Nodes))
	}
	if link.Nodes[0].NodeID != nodeA.NodeID {
		t.Errorf("Invalid node[0].id: %v", link.Nodes[0].NodeID)
	}
	if link.Nodes[0].AdapterNumber != nodeA.Ports[0].AdapterNumber {
		t.Errorf("Invalid node[0].adapterNumber: %v", link.Nodes[0].AdapterNumber)
	}
	if link.Nodes[0].PortNumber != nodeA.Ports[0].PortNumber {
		t.Errorf("Invalid node[0].portNumber: %v", link.Nodes[0].PortNumber)
	}
	if link.Nodes[0].Label.Text != "PortA" {
		t.Errorf("Invalid node[0].label.text: %v", link.Nodes[0].Label.Text)
	}
	if link.Nodes[0].Label.X != 5 {
		t.Errorf("Invalid node[0].label.x: %v", link.Nodes[0].Label.X)
	}
	if link.Nodes[0].Label.Y != 6 {
		t.Errorf("Invalid node[0].label.y: %v", link.Nodes[0].Label.Y)
	}
	if link.Nodes[0].Label.Rotation != 90 {
		t.Errorf("Invalid node[0].label.rotation: %v", link.Nodes[0].Label.Rotation)
	}
	if link.Nodes[0].Label.Style != "font-family: Verdana;" {
		t.Errorf("Invalid node[0].label.style: %v", link.Nodes[0].Label.Style)
	}
	if link.Nodes[1].NodeID != nodeB.NodeID {
		t.Errorf("Invalid node[1].id: %v", link.Nodes[1].NodeID)
	}
	if link.Nodes[1].AdapterNumber != nodeB.Ports[0].AdapterNumber {
		t.Errorf("Invalid node[1].adapterNumber: %v", link.Nodes[1].AdapterNumber)
	}
	if link.Nodes[1].PortNumber != nodeB.Ports[0].PortNumber {
		t.Errorf("Invalid node[1].portNumber: %v", link.Nodes[1].PortNumber)
	}
	if link.Nodes[1].Label.Text != "PortB" {
		t.Errorf("Invalid node[1].label.text: %v", link.Nodes[1].Label.Text)
	}
	if link.Nodes[1].Label.X != 7 {
		t.Errorf("Invalid node[1].label.x: %v", link.Nodes[1].Label.X)
	}
	if link.Nodes[1].Label.Y != 8 {
		t.Errorf("Invalid node[1].label.y: %v", link.Nodes[1].Label.Y)
	}
	if link.Nodes[1].Label.Rotation != 180 {
		t.Errorf("Invalid node[1].label.rotation: %v", link.Nodes[1].Label.Rotation)
	}
}
