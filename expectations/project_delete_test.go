package expectations

import (
	"gons3"
	"testing"
)

func TestDeleteProject(t *testing.T) {
	g := gons3.GNS3HTTPClient{}
	c := gons3.ProjectCreator{}

	c.SetName("TestDeleteProject")
	ci, err := gons3.CreateProject(g, c)
	if err != nil {
		t.Fatal(err)
	}

	err = gons3.DeleteProject(g, ci.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
}
