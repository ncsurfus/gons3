package gns3tests

import (
	"gons3"
)

var client = gons3.GNS3HTTPClient{}

func deleteProjectByName(g gons3.GNS3Client, name string) error {
	pjs, err := gons3.GetProjects(g)
	if err != nil {
		return err
	}
	for _, p := range pjs {
		if p.Name == name {
			gons3.DeleteProject(g, p.ProjectID)
		}
	}
	return nil
}
