package expectations

import (
	"gons3"
)

func deleteProjectByName(t gons3.Transport, name string) error {
	pjs, err := gons3.GetProjects(t)
	if err != nil {
		return err
	}
	for _, p := range pjs {
		if p.Name == name {
			gons3.DeleteProject(t, p.ProjectID)
		}
	}
	return nil
}
