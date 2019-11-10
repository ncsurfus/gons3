package expectations

import (
	"gons3"
	"testing"
)

func TestGetProject(t *testing.T) {
	g := gons3.GNS3HTTPClient{}
	c := gons3.ProjectCreator{}

	c.SetName("TestGetProject")
	ci, err := gons3.CreateProject(g, c)
	if err != nil {
		t.Fatal(err)
	}
	defer gons3.DeleteProject(g, ci.ProjectID)

	proj, err := gons3.GetProject(g, ci.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	if proj.Name != ci.Name {
		t.Fatal("Wrong Name!")
	}
}

func TestOpenCloseProject(t *testing.T) {
	g := gons3.GNS3HTTPClient{}
	c := gons3.ProjectCreator{}

	c.SetName("TestOpenCloseProject")
	ci, err := gons3.CreateProject(g, c)
	if err != nil {
		t.Fatal(err)
	}
	defer gons3.DeleteProject(g, ci.ProjectID)

	ci, err = gons3.OpenProject(g, ci.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	if !ci.IsOpened() {
		t.Fatal("Project was expected to be open, but wasn't.")
	}

	ci, err = gons3.CloseProject(g, ci.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	if ci.IsOpened() {
		t.Fatal("Project was expected to be closed, but wasn't.")
	}
}

func TestReadWriteProjectFile(t *testing.T) {
	g := gons3.GNS3HTTPClient{}
	c := gons3.ProjectCreator{}
	c.SetName("TestReadWriteProjectFile")
	ci, err := gons3.CreateProject(g, c)
	if err != nil {
		t.Fatal(err)
	}
	defer gons3.DeleteProject(g, ci.ProjectID)

	err = gons3.WriteProjectFile(g, ci.ProjectID, "testing", []byte("the test"))
	if err != nil {
		t.Fatal("Failed to write project file!")
	}

	data, err := gons3.ReadProjectFile(g, ci.ProjectID, "testing")
	sdata := string(data)
	if err != nil {
		t.Fatal("Failed to read project file!")
	}
	if sdata != "the test" {
		t.Fatal("Invalid data from project files!")
	}
}
