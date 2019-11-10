package expectations

import (
	"fmt"
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
	fmt.Print(ci)

	ci, err = gons3.CloseProject(g, ci.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(ci)
}
