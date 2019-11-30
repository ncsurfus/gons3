package gns3tests

import (
	"gons3"
	"testing"
)

func TestGetProject(t *testing.T) {
	projectBuilder := gons3.ProjectBuilder{}
	projectBuilder.SetName("TestGetProject")
	createdProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, createdProject.ProjectID)

	project, err := gons3.GetProject(client, createdProject.ProjectID)
	if err != nil {
		t.Fatalf("Error getting project: %v", err)
	}
	if project.Name != "TestGetProject" {
		t.Errorf("Expected name: %v, got %v", "TestGetProject", project.Name)
	}
}

func TestGetProjects(t *testing.T) {
	projectBuilderA := gons3.ProjectBuilder{}
	projectBuilderA.SetName("TestGetProjectsA")
	createdProjectA, err := gons3.CreateProject(client, projectBuilderA)
	if err != nil {
		t.Fatalf("Error creating project A: %v", err)
	}
	defer gons3.DeleteProject(client, createdProjectA.ProjectID)

	projectBuilderB := gons3.ProjectBuilder{}
	projectBuilderB.SetName("TestGetProjectsB")
	createdProjectB, err := gons3.CreateProject(client, projectBuilderB)
	if err != nil {
		t.Fatalf("Error creating project B: %v", err)
	}
	defer gons3.DeleteProject(client, createdProjectB.ProjectID)

	projects, err := gons3.GetProjects(client)
	if err != nil {
		t.Fatalf("Error getting projects: %v", err)
	}

	projectAFound, projectBFound := false, false
	for _, project := range projects {
		switch project.ProjectID {
		case createdProjectA.ProjectID:
			projectAFound = true
		case createdProjectB.ProjectID:
			projectBFound = true
		}
	}
	if !projectAFound {
		t.Errorf("projectA is missing!")
	}
	if !projectBFound {
		t.Errorf("projectB is missing!")
	}
}

func TestDeleteProject(t *testing.T) {
	c := gons3.ProjectBuilder{}
	c.SetName("TestDeleteProject")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}

	err = gons3.DeleteProject(client, ci.ProjectID)
	if err != nil {
		t.Fatalf("Error deleting project: %v", err)
	}
}

func TestOpenCloseProject(t *testing.T) {
	c := gons3.ProjectBuilder{}
	c.SetName("TestOpenCloseProject")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, ci.ProjectID)

	ci, err = gons3.OpenProject(client, ci.ProjectID)
	if err != nil {
		t.Fatalf("Error opening project: %v", err)
	}
	if !ci.IsOpened() {
		t.Errorf("Expected IsOpened(): %v, got %v", true, ci.IsOpened())
	}

	ci, err = gons3.CloseProject(client, ci.ProjectID)
	if err != nil {
		t.Fatalf("Error closing project: %v", err)
	}
	if ci.IsOpened() {
		t.Errorf("Expected IsOpened(): %v, got %v", false, ci.IsOpened())
	}
}

func TestReadWriteProjectFile(t *testing.T) {
	c := gons3.ProjectBuilder{}
	c.SetName("TestReadWriteProjectFile")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, ci.ProjectID)

	err = gons3.WriteProjectFile(client, ci.ProjectID, "testing", []byte("the test"))
	if err != nil {
		t.Fatalf("Error writing project file: %v", err)
	}

	data, err := gons3.ReadProjectFile(client, ci.ProjectID, "testing")
	sdata := string(data)
	if err != nil {
		t.Fatalf("Error reading project file: %v", err)
	}
	if sdata != "the test" {
		t.Errorf("Expected data: %v, got %v", "the test", sdata)
	}
}
