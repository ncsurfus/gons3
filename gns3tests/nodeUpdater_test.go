package gns3tests

import (
	"gons3"
	"testing"
)

func TestUpdateNode_A(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestUpdateNodeA")
	proj, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nodeBuilder := gons3.NewNodeBuilder("TheNode_A", "vpcs")
	newNode, err := gons3.CreateNode(client, proj.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, proj.ProjectID, newNode.ProjectID)

	nodeUpdater := gons3.NodeUpdater{}
	nodeUpdater.SetName("TheNode")
	nodeUpdater.SetConsoleType("none")
	nodeUpdater.SetConsoleAutoStart(true)
	nodeUpdater.SetX(5)
	nodeUpdater.SetY(6)
	nodeUpdater.SetZ(7)
	nodeUpdater.SetLocked(true)
	nodeUpdater.SetPortNameFormat("port-{0}")
	nodeUpdater.SetPortSegmentSize(1)
	nodeUpdater.SetFirstPortName("Mgmt0")
	// If the node name is changed, label.x gets centered, so no use in setting it.
	nodeUpdater.SetLabelY(15)
	nodeUpdater.SetLabelRotation(90)
	nodeUpdater.SetLabelStyle("font-family: TypeWriter;")
	node, err := gons3.UpdateNode(client, proj.ProjectID, newNode.NodeID, nodeUpdater)
	if err != nil {
		t.Fatalf("Error updating node: %v", err)
	}

	if node.Name != "TheNode" {
		t.Errorf("Expected name: %v, got %v", "TheNode", node.Name)
	}
	if node.ConsoleType != "none" {
		t.Errorf("Expected consoleType: %v, got %v", "none", node.ConsoleType)
	}
	if node.ConsoleAutoStart != true {
		t.Errorf("Expected consoleAutoStart: %v, got %v", true, node.ConsoleAutoStart)
	}
	if node.X != 5 {
		t.Errorf("Expected x: %v, got %v", 5, node.X)
	}
	if node.Y != 6 {
		t.Errorf("Expected y: %v, got %v", 6, node.Y)
	}
	if node.Z != 7 {
		t.Errorf("Expected z: %v, got %v", 7, node.Z)
	}
	if node.Label.Style != "font-family: TypeWriter;" {
		t.Errorf("Expected label.style: %v, got %v", "font-family: TypeWriter;", node.Label.Style)
	}
	if node.Label.Y != 15 {
		t.Errorf("Expected label.y: %v, got %v", 15, node.Label.Y)
	}
	if node.Label.Rotation != 90 {
		t.Errorf("Expected label.rotation: %v, got %v", 90, node.Label.Rotation)
	}
	if node.Locked != true {
		t.Errorf("Expected locked: %v, got %v", true, node.Locked)
	}
	if node.PortNameFormat != "port-{0}" {
		t.Errorf("Expected portNameFormat: %v, got %v", "port-{0}", node.PortNameFormat)
	}
	if node.PortSegmentSize != 1 {
		t.Errorf("Expected portSegmentSize: %v, got %v", 1, node.PortSegmentSize)
	}
	if node.FirstPortName != "Mgmt0" {
		t.Errorf("Expected firstPortName: %v, got %v", "Mgmt0", node.FirstPortName)
	}
}

func TestUpdateNode_B(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestUpdateNode_B")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilder := gons3.NewNodeBuilder("TheNode_B", "ethernet_hub")
	newNode, err := gons3.CreateNode(client, project.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, project.ProjectID, newNode.ProjectID)

	nodeUpdater := gons3.NodeUpdater{}
	nodeUpdater.SetConsoleType("none")
	nodeUpdater.SetConsoleAutoStart(false)
	nodeUpdater.SetX(8)
	nodeUpdater.SetY(9)
	nodeUpdater.SetZ(10)
	nodeUpdater.SetLocked(false)
	nodeUpdater.SetPortNameFormat("ports-{0}")
	nodeUpdater.SetPortSegmentSize(2)
	nodeUpdater.SetFirstPortName("Mgmt1")
	nodeUpdater.SetLabelX(20)
	nodeUpdater.SetLabelY(30)
	nodeUpdater.SetLabelRotation(180)
	nodeUpdater.SetLabelStyle("font-family: Verdana;")

	node, err := gons3.UpdateNode(client, project.ProjectID, newNode.NodeID, nodeUpdater)
	if err != nil {
		t.Fatalf("Error updating node: %v", err)
	}

	if node.Name != "TheNode_B" {
		t.Errorf("Expected name: %v, got %v", "TheNode_B", node.Name)
	}
	if node.ConsoleType != "none" {
		t.Errorf("Expected consoleType: %v, got %v", "none", node.ConsoleType)
	}
	if node.ConsoleAutoStart != false {
		t.Errorf("Expected consoleAutoStart: %v, got %v", false, node.ConsoleAutoStart)
	}
	if node.X != 8 {
		t.Errorf("Expected x: %v, got %v", 8, node.X)
	}
	if node.Y != 9 {
		t.Errorf("Expected y: %v, got %v", 9, node.Y)
	}
	if node.Z != 10 {
		t.Errorf("Expected z: %v, got %v", 10, node.Z)
	}
	if node.Label.Style != "font-family: Verdana;" {
		t.Errorf("Expected label.style: %v, got %v", "font-family: Verdana;", node.Label.Style)
	}
	if node.Label.X != 20 {
		t.Errorf("Expected label.x: %v, got %v", 20, node.Label.X)
	}
	if node.Label.Y != 30 {
		t.Errorf("Expected label.y: %v, got %v", 30, node.Label.Y)
	}
	if node.Label.Rotation != 180 {
		t.Errorf("Expected label.rotation: %v, got %v", 180, node.Label.Rotation)
	}
	if node.Locked != false {
		t.Errorf("Expected locked: %v, got %v", false, node.Locked)
	}
	if node.PortNameFormat != "ports-{0}" {
		t.Errorf("Expected portNameFormat: %v, got %v", "ports-{0}", node.PortNameFormat)
	}
	if node.PortSegmentSize != 2 {
		t.Errorf("Expected portSegmentSize: %v, got %v", 2, node.PortSegmentSize)
	}
	if node.FirstPortName != "Mgmt1" {
		t.Errorf("Expected firstPortName: %v, got %v", "Mgmt1", node.FirstPortName)
	}
}
