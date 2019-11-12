package gns3tests

import (
	"gons3"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	pc := gons3.ProjectCreator{}
	pc.SetName("TestDeleteNode")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeCreator{}
	nc.SetName("TheNode")
	nc.SetNodeType("vpcs")
	nc.SetLocalComputeID()
	node, err := gons3.CreateNode(client, proj.ProjectID, nc)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	err = gons3.DeleteNode(client, proj.ProjectID, node.NodeID)
	if err != nil {
		t.Fatalf("Error deleting node: %v", err)
	}
}

func TestGetNode(t *testing.T) {
	pc := gons3.ProjectCreator{}
	pc.SetName("TestGetNode")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeCreator{}
	nc.SetName("TheNode")
	nc.SetNodeType("vpcs")
	nc.SetLocalComputeID()
	nc.SetConsoleType("telnet")
	nc.SetConsoleAutoStart(true)
	nc.SetX(5)
	nc.SetY(6)
	nc.SetZ(7)
	nc.SetLocked(true)
	nc.SetPortNameFormat("port-{0}")
	nc.SetPortSegmentSize(1)
	nc.SetFirstPortName("Mgmt0")

	lc := gons3.LabelCreator{}
	lc.SetX(10)
	lc.SetY(15)
	lc.SetRotation(90)
	lc.SetStyle("font-family: TypeWriter;")
	nc.SetLabel(lc)

	lc := gons3.CustomAdapterCreator{}
	lc.


	node, err := gons3.CreateNode(client, proj.ProjectID, nc)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, proj.ProjectID, node.ProjectID)

	if node.Name != "TheNode" {
		t.Errorf("Expected name: %v, got %v", "TestCreateA", node.Name)
	}
	if node.NodeType != "vpcs" {
		t.Errorf("Expected nodeType: %v, got %v", "vpcs", node.NodeType)
	}
	if node.ConsoleType != "telnet" {
		t.Errorf("Expected consoleType: %v, got %v", "telnet", node.ConsoleType)
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
	if node.Label.X != 10 {
		t.Errorf("Expected label.x: %v, got %v", 10, node.Label.X)
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
