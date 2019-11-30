package gns3tests

import (
	"gons3"
	"testing"
)

func TestGetLink(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestGetLink")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		fatalAssert(t, "CreateProject error", nil, err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NewNodeBuilder("TheNodeA", "vpcs")
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		fatalAssert(t, "CreateNode(A) error", nil, err)
	}

	nodeBuilderB := gons3.NewNodeBuilder("TheNodeB", "vpcs")
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		fatalAssert(t, "CreateNode(B) error", nil, err)
	}

	linkBuilder := gons3.NewLinkBuilder(nodeA.Ports[0], nodeB.Ports[0])
	createdLink, err := gons3.CreateLink(client, project.ProjectID, linkBuilder)
	if err != nil {
		fatalAssert(t, "CreateLink error", nil, err)
	}

	link, err := gons3.GetLink(client, project.ProjectID, createdLink.LinkID)
	if err != nil {
		fatalAssert(t, "GetLink error", nil, err)
	}
	fatalAssert(t, "len(link.Nodes)", 2, len(link.Nodes))
	errorAssert(t, "link.Nodes[0].NodeID", nodeA.NodeID, link.Nodes[0].NodeID)
	errorAssert(t, "link.Nodes[0].AdapterNumber", nodeA.Ports[0].AdapterNumber, link.Nodes[0].AdapterNumber)
	errorAssert(t, "link.Nodes[0].PortNumber", nodeA.Ports[0].PortNumber, link.Nodes[0].PortNumber)
	errorAssert(t, "link.Nodes[1].NodeID", nodeB.NodeID, link.Nodes[1].NodeID)
	errorAssert(t, "link.Nodes[1].AdapterNumber", nodeB.Ports[0].AdapterNumber, link.Nodes[1].AdapterNumber)
	errorAssert(t, "link.Nodes[1].PortNumber", nodeB.Ports[0].PortNumber, link.Nodes[1].PortNumber)
}

func TestGetLinks(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestGetLinks")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NewNodeBuilder("TheNodeA", "ethernet_switch")
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeBuilderB := gons3.NewNodeBuilder("TheNodeB", "ethernet_switch")
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	linkBuilderA := gons3.NewLinkBuilder(nodeA.Ports[0], nodeB.Ports[0])
	createdLinkA, err := gons3.CreateLink(client, project.ProjectID, linkBuilderA)
	if err != nil {
		t.Fatalf("Error creating linkA: %v", err)
	}

	linkBuilderB := gons3.NewLinkBuilder(nodeA.Ports[1], nodeB.Ports[1])
	createdLinkB, err := gons3.CreateLink(client, project.ProjectID, linkBuilderB)
	if err != nil {
		t.Fatalf("Error creating linkB: %v", err)
	}

	link, err := gons3.GetLinks(client, project.ProjectID)
	if err != nil {
		t.Fatalf("Error getting link: %v", err)
	}
	if len(link) != 2 {
		t.Fatalf("Invalid link count: %v", len(link))
	}
	if link[0].LinkID != createdLinkA.LinkID {
		t.Errorf("Invalid link.0.linkID: %v", createdLinkA.LinkID)
	}
	if link[1].LinkID != createdLinkB.LinkID {
		t.Errorf("Invalid link.1.linkID: %v", createdLinkB.LinkID)
	}
}

func TestGetNodeLinks(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestGetNodeLinks")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NewNodeBuilder("TheNodeA", "ethernet_switch")
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeBuilderB := gons3.NewNodeBuilder("TheNodeB", "ethernet_switch")
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	linkBuilderA := gons3.NewLinkBuilder(nodeA.Ports[0], nodeB.Ports[0])
	createdLinkA, err := gons3.CreateLink(client, project.ProjectID, linkBuilderA)
	if err != nil {
		t.Fatalf("Error creating linkA: %v", err)
	}

	linkBuilderB := gons3.NewLinkBuilder(nodeA.Ports[1], nodeB.Ports[1])
	createdLinkB, err := gons3.CreateLink(client, project.ProjectID, linkBuilderB)
	if err != nil {
		t.Fatalf("Error creating linkB: %v", err)
	}

	link, err := gons3.GetNodeLinks(client, project.ProjectID, nodeA.NodeID)
	if err != nil {
		t.Fatalf("Error getting node links: %v", err)
	}
	if len(link) != 2 {
		t.Fatalf("Invalid link count: %v", len(link))
	}
	if (link[0].LinkID != createdLinkA.LinkID && link[1].LinkID != createdLinkB.LinkID) &&
		(link[0].LinkID != createdLinkB.LinkID && link[1].LinkID != createdLinkA.LinkID) {
		t.Errorf("Invalid linkIDs: %v, %v", createdLinkA.LinkID, createdLinkB.LinkID)
	}
}

func TestDeleteLink(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestDeleteLink")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	nodeBuilderA := gons3.NewNodeBuilder("TheNodeA", "vpcs")
	nodeA, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderA)
	if err != nil {
		t.Fatalf("Error creating nodeA: %v", err)
	}

	nodeBuilderB := gons3.NewNodeBuilder("TheNodeB", "vpcs")
	nodeB, err := gons3.CreateNode(client, project.ProjectID, nodeBuilderB)
	if err != nil {
		t.Fatalf("Error creating nodeB: %v", err)
	}

	linkBuilder := gons3.NewLinkBuilder(nodeA.Ports[0], nodeB.Ports[0])
	createdLink, err := gons3.CreateLink(client, project.ProjectID, linkBuilder)
	if err != nil {
		t.Fatalf("Error creating link: %v", err)
	}

	err = gons3.DeleteLink(client, project.ProjectID, createdLink.LinkID)
	if err != nil {
		t.Fatalf("Error deleting link: %v", err)
	}

	_, err = gons3.GetLink(client, project.ProjectID, createdLink.LinkID)
	if err == nil {
		t.Fatalf("Got link after deleting link: %v", err)
	}
}
