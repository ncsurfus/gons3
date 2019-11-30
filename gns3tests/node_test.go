package gns3tests

import (
	"gons3"
	"testing"
	"time"
)

func TestGetNode(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestGetNode")
	createdProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, createdProject.ProjectID)

	nodeBuilder := gons3.NewNodeBuilder("TheNode", "vpcs")
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
	projectBuilder := gons3.NewProjectBuilder("TestGetNodes")
	createdProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, createdProject.ProjectID)

	nodeBuilderA := gons3.NewNodeBuilder("TheNodeA", "vpcs")
	_, err = gons3.CreateNode(client, createdProject.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	nodeBuilderB := gons3.NewNodeBuilder("TheNodeB", "vpcs")
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
	projectBuilder := gons3.NewProjectBuilder("TestDeleteNode")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilder := gons3.NewNodeBuilder("TheNode", "vpcs")
	node, err := gons3.CreateNode(client, project.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	err = gons3.DeleteNode(client, project.ProjectID, node.NodeID)
	if err != nil {
		t.Fatalf("Error deleting node: %v", err)
	}
}

func TestStartStopNode(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestStartStopNode")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilder := gons3.NewNodeBuilder("TheNode", "vpcs")
	node, err := gons3.CreateNode(client, project.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, project.ProjectID, node.ProjectID)

	node, err = gons3.StartNode(client, project.ProjectID, node.NodeID)
	if err != nil {
		t.Fatalf("Error starting node: %v", err)
	}
	// Wait 1 second, 10 times to see if node started
	for i := 0; i != 10; i++ {
		node, err = gons3.GetNode(client, project.ProjectID, node.NodeID)
		if node.IsStarted() {
			break
		}
		time.Sleep(time.Second * 1)
	}
	if !node.IsStarted() {
		t.Fatalf("Node did not start: %v", node.Status)
	}

	node, err = gons3.StopNode(client, project.ProjectID, node.NodeID)
	if err != nil {
		t.Fatalf("Error stopping node: %v", err)
	}
	// Wait 1 second, 10 times to see if node started
	for i := 0; i != 10; i++ {
		node, err = gons3.GetNode(client, project.ProjectID, node.NodeID)
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
	projectBuilder := gons3.NewProjectBuilder("TestReadWriteNodeFile")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilder := gons3.NewNodeBuilder("TheNode", "vpcs")
	node, err := gons3.CreateNode(client, project.ProjectID, nodeBuilder)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}
	defer gons3.DeleteNode(client, project.ProjectID, node.ProjectID)

	err = gons3.WriteNodeFile(client, project.ProjectID, node.NodeID, "test.txt", []byte("Test123"))
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	b, err := gons3.ReadNodeFile(client, project.ProjectID, node.NodeID, "test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}
	if string(b) != "Test123" {
		t.Fatalf("Invalid file data : %v", string(b))
	}
}
