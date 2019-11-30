package gns3tests

import (
	"gons3"
	"testing"
)

func TestGetProject(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestGetProject")
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
	projectBuilderA := gons3.NewProjectBuilder("TestGetProjectsA")
	createdProjectA, err := gons3.CreateProject(client, projectBuilderA)
	if err != nil {
		t.Fatalf("Error creating project A: %v", err)
	}
	defer gons3.DeleteProject(client, createdProjectA.ProjectID)

	projectBuilderB := gons3.NewProjectBuilder("TestGetProjectsB")
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
	projectBuilder := gons3.NewProjectBuilder("TestDeleteProject")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}

	err = gons3.DeleteProject(client, project.ProjectID)
	if err != nil {
		t.Fatalf("Error deleting project: %v", err)
	}
}

func TestOpenCloseProject(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestOpenCloseProject")
	createdProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, createdProject.ProjectID)

	openedProject, err := gons3.OpenProject(client, createdProject.ProjectID)
	if err != nil {
		t.Fatalf("Error opening project: %v", err)
	}
	if !openedProject.IsOpened() {
		t.Errorf("Expected IsOpened(): %v, got %v", true, openedProject.IsOpened())
	}

	closedProject, err := gons3.CloseProject(client, createdProject.ProjectID)
	if err != nil {
		t.Fatalf("Error closing project: %v", err)
	}
	if closedProject.IsOpened() {
		t.Errorf("Expected IsOpened(): %v, got %v", false, closedProject.IsOpened())
	}
}

func TestReadWriteProjectFile(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestReadWriteProjectFile")
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	err = gons3.WriteProjectFile(client, project.ProjectID, "testing", []byte("the test"))
	if err != nil {
		t.Fatalf("Error writing project file: %v", err)
	}

	data, err := gons3.ReadProjectFile(client, project.ProjectID, "testing")
	sdata := string(data)
	if err != nil {
		t.Fatalf("Error reading project file: %v", err)
	}
	if sdata != "the test" {
		t.Errorf("Expected data: %v, got %v", "the test", sdata)
	}
}
