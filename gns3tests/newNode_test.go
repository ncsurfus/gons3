package gns3tests

import (
	"gons3"
	"testing"
)

func TestCreateNode_A(t *testing.T) {
	pc := gons3.ProjectBuilder{}
	pc.SetName("TestCreateNode_A")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeBuilder{}
	nc.SetName("TheNode_A")
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
	nc.SetLabelX(10)
	nc.SetLabelY(15)
	nc.SetLabelRotation(90)
	nc.SetLabelStyle("font-family: TypeWriter;")

	node, err := gons3.CreateNode(client, proj.ProjectID, nc)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, proj.ProjectID, node.ProjectID)

	if node.Name != "TheNode_A" {
		t.Errorf("Expected name: %v, got %v", "TheNode_A", node.Name)
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

func TestCreateNode_B(t *testing.T) {
	pc := gons3.ProjectBuilder{}
	pc.SetName("TestCreateNode_B")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeBuilder{}
	nc.SetName("TheNode_B")
	nc.SetNodeType("ethernet_hub")
	nc.SetLocalComputeID()
	nc.SetConsoleType("none")
	nc.SetConsoleAutoStart(false)
	nc.SetX(8)
	nc.SetY(9)
	nc.SetZ(10)
	nc.SetLocked(false)
	nc.SetPortNameFormat("ports-{0}")
	nc.SetPortSegmentSize(2)
	nc.SetFirstPortName("Mgmt1")
	nc.SetLabelX(20)
	nc.SetLabelY(30)
	nc.SetLabelRotation(180)
	nc.SetLabelStyle("font-family: Verdana;")

	node, err := gons3.CreateNode(client, proj.ProjectID, nc)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, proj.ProjectID, node.ProjectID)

	if node.Name != "TheNode_B" {
		t.Errorf("Expected name: %v, got %v", "TheNode_B", node.Name)
	}
	if node.NodeType != "ethernet_hub" {
		t.Errorf("Expected nodeType: %v, got %v", "ethernet_hub", node.NodeType)
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
