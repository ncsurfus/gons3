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
