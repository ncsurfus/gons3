package gns3tests

import (
	"gons3"
	"testing"
)

func TestCreateNode_A(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestCreateNodeA")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilder := gons3.NodeBuilder{}
	nodeBuilder.SetName("TheNode_A")
	nodeBuilder.SetNodeType("vpcs")
	nodeBuilder.SetLocalComputeID()
	nodeBuilder.SetConsoleType("telnet")
	nodeBuilder.SetConsoleAutoStart(true)
	nodeBuilder.SetX(5)
	nodeBuilder.SetY(6)
	nodeBuilder.SetZ(7)
	nodeBuilder.SetLocked(true)
	nodeBuilder.SetPortNameFormat("port-{0}")
	nodeBuilder.SetPortSegmentSize(1)
	nodeBuilder.SetFirstPortName("Mgmt0")
	nodeBuilder.SetLabelX(10)
	nodeBuilder.SetLabelY(15)
	nodeBuilder.SetLabelRotation(90)
	nodeBuilder.SetLabelStyle("font-family: TypeWriter;")

	node, err := gons3.CreateNode(client, project.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, project.ProjectID, node.ProjectID)

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
	projectBuilder := gons3.NewProjectBuilder("TestCreateNodeB")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilder := gons3.NodeBuilder{}
	nodeBuilder.SetName("TheNode_B")
	nodeBuilder.SetNodeType("ethernet_hub")
	nodeBuilder.SetLocalComputeID()
	nodeBuilder.SetConsoleType("none")
	nodeBuilder.SetConsoleAutoStart(false)
	nodeBuilder.SetX(8)
	nodeBuilder.SetY(9)
	nodeBuilder.SetZ(10)
	nodeBuilder.SetLocked(false)
	nodeBuilder.SetPortNameFormat("ports-{0}")
	nodeBuilder.SetPortSegmentSize(2)
	nodeBuilder.SetFirstPortName("Mgmt1")
	nodeBuilder.SetLabelX(20)
	nodeBuilder.SetLabelY(30)
	nodeBuilder.SetLabelRotation(180)
	nodeBuilder.SetLabelStyle("font-family: Verdana;")

	node, err := gons3.CreateNode(client, project.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, project.ProjectID, node.ProjectID)

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
