package gns3tests

import (
	"gons3"
	"testing"
)

func TestCreateLinkA(t *testing.T) {
	projectCreate := gons3.ProjectCreate{}
	projectCreate.SetName("TestCreateLinkA")
	project, err := gons3.CreateProject(client, projectCreate)
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
	linkNodeCreateA.SetLabelX(5)
	linkNodeCreateA.SetLabelY(6)
	linkNodeCreateA.SetLabelRotation(90)
	linkNodeCreateA.SetLabelStyle("font-family: Verdana;")
	linkNodeCreateA.SetLabelText("PortA")

	linkNodeCreateB := gons3.LinkNodeCreate{
		NodeID:        nodeB.NodeID,
		AdapterNumber: nodeB.Ports[0].AdapterNumber,
		PortNumber:    nodeB.Ports[0].PortNumber,
	}
	linkNodeCreateB.SetLabelX(7)
	linkNodeCreateB.SetLabelY(8)
	linkNodeCreateB.SetLabelRotation(180)
	linkNodeCreateB.SetLabelText("PortB")

	linkCreate := gons3.LinkCreate{}
	linkCreate.SetLinkType("ethernet")
	linkCreate.SetSuspend(false)
	linkCreate.SetNodes([]gons3.LinkNodeCreate{linkNodeCreateA, linkNodeCreateB})
	link, err := gons3.CreateLink(client, project.ProjectID, linkCreate)
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
