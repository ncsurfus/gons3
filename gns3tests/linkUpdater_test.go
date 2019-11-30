package gns3tests

import (
	"gons3"
	"testing"
)

func TestUpdateLinkA(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestCreateLinkA")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NewNodeBuilder("TheNodeA", "ethernet_hub")
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeBuilderB := gons3.NewNodeBuilder("TheNodeB", "ethernet_hub")
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	linkBuilder := gons3.NewLinkBuilder(nodeA.Ports[1], nodeB.Ports[1])
	newLink, err := gons3.CreateLink(client, project.ProjectID, linkBuilder)
	if err != nil {
		t.Fatalf("Error creating link: %v", err)
	}

	linkNodeUpdaterA := gons3.LinkNodeUpdater{}
	linkNodeUpdaterA.SetNodeID(nodeA.NodeID)
	linkNodeUpdaterA.SetLabelX(5)
	linkNodeUpdaterA.SetLabelY(6)
	linkNodeUpdaterA.SetLabelRotation(90)
	linkNodeUpdaterA.SetLabelStyle("font-family: Verdana;")
	linkNodeUpdaterA.SetLabelText("PortA")

	linkNodeUpdaterB := gons3.LinkNodeUpdater{}
	linkNodeUpdaterB.SetNodeID(nodeB.NodeID)
	linkNodeUpdaterB.SetLabelX(7)
	linkNodeUpdaterB.SetLabelY(8)
	linkNodeUpdaterB.SetLabelRotation(180)
	linkNodeUpdaterB.SetLabelText("PortB")

	linkUpdater := gons3.LinkUpdater{}
	linkUpdater.SetSuspend(true)
	linkUpdater.SetNodes(linkNodeUpdaterA, linkNodeUpdaterB)
	link, err := gons3.UpdateLink(client, project.ProjectID, newLink.LinkID, linkUpdater)
	if err != nil {
		t.Fatalf("Error updating link: %v", err)
	}

	if link.LinkType != "ethernet" {
		t.Errorf("Invalid link type: %v", link.LinkType)
	}
	if link.Suspended != true {
		t.Errorf("Invalid suspend: %v", link.Suspended)
	}
	if len(link.Nodes) != 2 {
		t.Fatalf("Invalid node count: %v", len(link.Nodes))
	}
	if link.Nodes[0].NodeID != nodeA.NodeID {
		t.Errorf("Invalid node[0].id: %v", link.Nodes[0].NodeID)
	}
	if link.Nodes[0].AdapterNumber != nodeA.Ports[1].AdapterNumber {
		t.Errorf("Invalid node[0].adapterNumber: %v", link.Nodes[0].AdapterNumber)
	}
	if link.Nodes[0].PortNumber != nodeA.Ports[1].PortNumber {
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
	if link.Nodes[1].AdapterNumber != nodeB.Ports[1].AdapterNumber {
		t.Errorf("Invalid node[1].adapterNumber: %v", link.Nodes[1].AdapterNumber)
	}
	if link.Nodes[1].PortNumber != nodeB.Ports[1].PortNumber {
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
