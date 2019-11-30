package gns3tests

import (
	"gons3"
	"testing"
	"time"
)

func TestGetNode(t *testing.T) {
	projectBuilder := gons3.ProjectBuilder{}
	projectBuilder.SetName("TestGetNode")
	createdProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, createdProject.ProjectID)

	nodeBuilder := gons3.NodeBuilder{}
	nodeBuilder.SetName("TheNode")
	nodeBuilder.SetNodeType("vpcs")
	nodeBuilder.SetLocalComputeID()
	createdNode, err := gons3.CreateNode(client, createdProject.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	node, err := gons3.GetNode(client, createdProject.ProjectID, createdNode.NodeID)
	if err != nil {
		t.Fatalf("Error getting node: %v", err)
	}
	if node.Name != "TheNode" {
		t.Errorf("Expected name: %v, got %v", "TestGetNode", node.Name)
	}
}

func TestGetNodes(t *testing.T) {
	projectBuilder := gons3.ProjectBuilder{}
	projectBuilder.SetName("TestGetNodes")
	createdProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, createdProject.ProjectID)

	nodeBuilderA := gons3.NodeBuilder{}
	nodeBuilderA.SetName("TheNodeA")
	nodeBuilderA.SetNodeType("vpcs")
	nodeBuilderA.SetLocalComputeID()
	_, err = gons3.CreateNode(client, createdProject.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	nodeBuilderB := gons3.NodeBuilder{}
	nodeBuilderB.SetName("TheNodeA")
	nodeBuilderB.SetNodeType("vpcs")
	nodeBuilderB.SetLocalComputeID()
	_, err = gons3.CreateNode(client, createdProject.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	nodes, err := gons3.GetNodes(client, createdProject.ProjectID)
	if err != nil {
		t.Fatalf("Error getting nodes: %v", err)
	}
	if len(nodes) != 2 {
		t.Fatalf("Expected node length: %v, got %v", 2, len(nodes))
	}
}

func TestDeleteNode(t *testing.T) {
	pc := gons3.ProjectBuilder{}
	pc.SetName("TestDeleteNode")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeBuilder{}
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

func TestStartStopNode(t *testing.T) {
	pc := gons3.ProjectBuilder{}
	pc.SetName("TestStartStopNode")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeBuilder{}
	nc.SetName("TheNode")
	nc.SetNodeType("vpcs")
	nc.SetLocalComputeID()
	node, err := gons3.CreateNode(client, proj.ProjectID, nc)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, proj.ProjectID, node.ProjectID)

	node, err = gons3.StartNode(client, proj.ProjectID, node.NodeID)
	if err != nil {
		t.Fatalf("Error starting node: %v", err)
	}
	// Wait 1 second, 10 times to see if node started
	for i := 0; i != 10; i++ {
		node, err = gons3.GetNode(client, proj.ProjectID, node.NodeID)
		if node.IsStarted() {
			break
		}
		time.Sleep(time.Second * 1)
	}
	if !node.IsStarted() {
		t.Fatalf("Node did not start: %v", node.Status)
	}

	node, err = gons3.StopNode(client, proj.ProjectID, node.NodeID)
	if err != nil {
		t.Fatalf("Error stopping node: %v", err)
	}
	// Wait 1 second, 10 times to see if node started
	for i := 0; i != 10; i++ {
		node, err = gons3.GetNode(client, proj.ProjectID, node.NodeID)
		if node.IsStopped() {
			break
		}
		time.Sleep(time.Second * 1)
	}
	if !node.IsStopped() {
		t.Fatalf("Node did not stop: %v", node.Status)
	}
}

func TestReadWriteNodeFile(t *testing.T) {
	pc := gons3.ProjectBuilder{}
	pc.SetName("TestWriteNodeFile")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeBuilder{}
	nc.SetName("TheNode")
	nc.SetNodeType("vpcs")
	nc.SetLocalComputeID()
	node, err := gons3.CreateNode(client, proj.ProjectID, nc)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, proj.ProjectID, node.ProjectID)

	err = gons3.WriteNodeFile(client, proj.ProjectID, node.NodeID, "test.txt", []byte("Test123"))
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	b, err := gons3.ReadNodeFile(client, proj.ProjectID, node.NodeID, "test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}
	if string(b) != "Test123" {
		t.Fatalf("Invalid file data : %v", string(b))
	}
}
