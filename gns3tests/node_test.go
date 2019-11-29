package gns3tests

import (
	"gons3"
	"testing"
	"time"
)

func TestGetNode(t *testing.T) {
	projectCreate := gons3.ProjectCreate{}
	projectCreate.SetName("TestGetNode")
	newProject, err := gons3.CreateProject(client, projectCreate)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, newProject.ProjectID)

	nodeCreate := gons3.NodeCreate{}
	nodeCreate.SetName("TheNode")
	nodeCreate.SetNodeType("vpcs")
	nodeCreate.SetLocalComputeID()
	newNode, err := gons3.CreateNode(client, newProject.ProjectID, nodeCreate)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	node, err := gons3.GetNode(client, newProject.ProjectID, newNode.NodeID)
	if err != nil {
		t.Fatalf("Error getting node: %v", err)
	}
	if node.Name != "TheNode" {
		t.Errorf("Expected name: %v, got %v", "TestGetNode", node.Name)
	}
}

func TestGetNodes(t *testing.T) {
	projectCreate := gons3.ProjectCreate{}
	projectCreate.SetName("TestGetNodes")
	newProject, err := gons3.CreateProject(client, projectCreate)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, newProject.ProjectID)

	nodeCreateA := gons3.NodeCreate{}
	nodeCreateA.SetName("TheNodeA")
	nodeCreateA.SetNodeType("vpcs")
	nodeCreateA.SetLocalComputeID()
	_, err = gons3.CreateNode(client, newProject.ProjectID, nodeCreateA)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	nodeCreateB := gons3.NodeCreate{}
	nodeCreateB.SetName("TheNodeA")
	nodeCreateB.SetNodeType("vpcs")
	nodeCreateB.SetLocalComputeID()
	_, err = gons3.CreateNode(client, newProject.ProjectID, nodeCreateB)
	if err != nil {
		t.Fatalf("Error creating node: %v", err)
	}

	nodes, err := gons3.GetNodes(client, newProject.ProjectID)
	if err != nil {
		t.Fatalf("Error getting nodes: %v", err)
	}
	if len(nodes) != 2 {
		t.Fatalf("Expected node length: %v, got %v", 2, len(nodes))
	}
}

func TestDeleteNode(t *testing.T) {
	pc := gons3.ProjectCreate{}
	pc.SetName("TestDeleteNode")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeCreate{}
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
	pc := gons3.ProjectCreate{}
	pc.SetName("TestStartStopNode")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeCreate{}
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
	pc := gons3.ProjectCreate{}
	pc.SetName("TestWriteNodeFile")
	proj, err := gons3.CreateProject(client, pc)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, proj.ProjectID)

	nc := gons3.NodeCreate{}
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
