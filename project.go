// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/schemas/project.py

package gons3

import (
	"errors"
	"fmt"
	"net/url"
)

// ProjectSupplier models a GNS3 Project's Supplier
type ProjectSupplier struct {
	Logo string `json:"logo"`
	URL  string `json:"url"`
}

// ProjectVariables models a GNS3 Project's Variables
type ProjectVariables struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Project models an instance of a GNS3 project.
type Project struct {
	Name                string              `json:"name"`
	ProjectID           string              `json:"project_id"`
	Path                string              `json:"path"`
	Filename            string              `json:"filename"`
	Status              string              `json:"status"`
	AutoClose           bool                `json:"auto_close"`
	AutoOpen            bool                `json:"auto_open"`
	AutoStart           bool                `json:"auto_start"`
	SceneHeight         int                 `json:"scene_height"`
	SceneWidth          int                 `json:"scene_width"`
	Zoom                int                 `json:"zoom"`
	ShowLayers          bool                `json:"show_layers"`
	SnapToGrid          bool                `json:"snap_to_grid"`
	ShowGrid            bool                `json:"show_grid"`
	GridSize            int                 `json:"grid_size"`
	ShowInterfaceLabels bool                `json:"show_interface_labels"`
	Supplier            *ProjectSupplier    `json:"Supplier"`
	Variables           *[]ProjectVariables `json:"Variables"`
}

// CreateProject creates a GNS3 project with the specified name.
func CreateProject(t Transport, p ProjectCreator) (Project, error) {
	if p.values == nil {
		return Project{}, errors.New("failed to create project: missing project name")
	}
	if name, ok := p.values["name"]; !ok || name == "" {
		return Project{}, errors.New("failed to create project: invalid project name")
	}

	u := "/v2/projects"
	i := Project{}
	_, err := t.Post(u, p.values, &i)
	if err != nil {
		return Project{}, fmt.Errorf("failed to create project: %v", err)
	}

	return i, nil
}

// UpdateProject creates a GNS3 project with the specified name.
func UpdateProject(t Transport, id string, p ProjectUpdater) (Project, error) {
	if p.values == nil {
		return Project{}, errors.New("failed to update project: nothing to update")
	}
	if id == "" {
		return Project{}, errors.New("failed to update project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id)
	i := Project{}
	_, err := t.Put(u, p.values, &i)
	if err != nil {
		return Project{}, fmt.Errorf("failed to update project: %v", err)
	}

	return i, nil
}

// DeleteProject deletes a GNS3 project instance with the specified id.
func DeleteProject(t Transport, id string) error {
	if id == "" {
		return errors.New("failed to delete project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id)
	_, err := t.Delete(u, nil)
	if err != nil {
		return fmt.Errorf("failed to delete project: %v", err)
	}

	return nil
}

// GetProject gets a GNS3 project instance with the specified id.
func GetProject(t Transport, id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("failed to get project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id)
	i := Project{}
	_, err := t.Get(u, &i)
	if err != nil {
		return Project{}, fmt.Errorf("failed to get project: %v", err)
	}

	return i, nil
}

// GetProjects gets all the GNS3 projects.
func GetProjects(t Transport) ([]Project, error) {
	u := "/v2/projects"
	i := []Project{}
	_, err := t.Get(u, &i)
	if err != nil {
		return []Project{}, fmt.Errorf("failed to get projects: %v", err)
	}

	return i, nil
}

// OpenProject opens the GNS3 project.
func OpenProject(t Transport, id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("failed to open project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id) + "/open"
	i := Project{}
	_, err := t.Post(u, nil, &i)
	if err != nil {
		return Project{}, fmt.Errorf("failed to open project: %v", err)
	}

	return i, nil
}

// CloseProject opens the GNS3 project.
func CloseProject(t Transport, id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("failed to close project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id) + "/close"
	i := Project{}
	_, err := t.Post(u, nil, &i)
	if err != nil {
		return Project{}, fmt.Errorf("failed to close project: %v", err)
	}

	return i, nil
}

// ReadProjectFile reads a GNS3 project's file.
func ReadProjectFile(t Transport, id string, path string) ([]byte, error) {
	if id == "" {
		return []byte{}, errors.New("failed to read project file: invalid project id")
	}
	if path == "" {
		return []byte{}, errors.New("failed to read project file: invalid path")
	}

	u := "/v2/projects/" + url.PathEscape(id) + "/files/" + path
	b := []byte{}
	_, err := t.Get(u, &b)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to close project: %v", err)
	}

	return b, nil
}

// WriteProjectFile writes a GNS3 project's file.
func WriteProjectFile(t Transport, id string, path string, data []byte) error {
	if id == "" {
		return errors.New("failed to write project file: invalid project id")
	}
	if path == "" {
		return errors.New("failed to write project file: invalid path")
	}

	u := "/v2/projects/" + url.PathEscape(id) + "/files/" + path
	_, err := t.Post(u, &data, nil)
	if err != nil {
		return fmt.Errorf("failed to close project: %v", err)
	}

	return nil
}
