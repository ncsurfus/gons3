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

// NewProject models a new GNS3 project.
type NewProject struct {
	Name                *string             `json:"name"`
	Path                *string             `json:"path,omitempty"`
	AutoClose           bool                `json:"auto_close,omitempty"`
	ProjectID           *string             `json:"project_id,omitempty"`
	SceneHeight         int                 `json:"scene_height,omitempty"`
	SceneWidth          int                 `json:"scene_width,omitempty"`
	Zoom                int                 `json:"zoom,omitempty"`
	ShowLayers          bool                `json:"show_layers,omitempty"`
	SnapToGrid          bool                `json:"snap_to_grid,omitempty"`
	ShowGrid            bool                `json:"show_grid,omitempty"`
	GridSize            int                 `json:"grid_size,omitempty"`
	DrawingGridSize     int                 `json:"drawing_grid_size,omitempty"`
	ShowInterfaceLabels bool                `json:"show_interface_labels,omitempty"`
	Supplier            *ProjectSupplier    `json:"Supplier,omitempty"`
	Variables           *[]ProjectVariables `json:"Variables,omitempty"`
}

// ProjectUpdate models an update to a GNS3 project.
type ProjectUpdate struct {
	Name                *string             `json:"name,omitempty"`
	Path                *string             `json:"path,omitempty"`
	AutoClose           bool                `json:"auto_close,omitempty"`
	AutoOpen            bool                `json:"auto_open,omitempty"`
	AutoStart           bool                `json:"auto_start,omitempty"`
	SceneHeight         int                 `json:"scene_height,omitempty"`
	SceneWidth          int                 `json:"scene_width,omitempty"`
	Zoom                int                 `json:"zoom,omitempty"`
	ShowLayers          bool                `json:"show_layers,omitempty"`
	SnapToGrid          bool                `json:"snap_to_grid,omitempty"`
	ShowGrid            bool                `json:"show_grid,omitempty"`
	GridSize            int                 `json:"grid_size,omitempty"`
	DrawingGridSize     int                 `json:"drawing_grid_size,omitempty"`
	ShowInterfaceLabels bool                `json:"show_interface_labels,omitempty"`
	Supplier            *ProjectSupplier    `json:"Supplier,omitempty"`
	Variables           *[]ProjectVariables `json:"Variables,omitempty"`
}

// ProjectInstance models an instance of a GNS3 project.
type ProjectInstance struct {
	Name                *string             `json:"name"`
	ProjectID           string              `json:"project_id"`
	Path                *string             `json:"path"`
	Filename            *string             `json:"filename"`
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
	DrawingGridSize     int                 `json:"drawing_grid_size"`
	ShowInterfaceLabels bool                `json:"show_interface_labels"`
	Supplier            *ProjectSupplier    `json:"Supplier"`
	Variables           *[]ProjectVariables `json:"Variables"`
}

// CreateProject creates a GNS3 project with the specified name.
func CreateProject(t Transport, p NewProject) (ProjectInstance, error) {
	if p.Name == nil || *p.Name == "" {
		return ProjectInstance{}, errors.New("failed to create project: invalid project name")
	}

	u := "/v2/projects"
	i := ProjectInstance{}
	_, err := t.Post(u, p, &i)
	if err != nil {
		return ProjectInstance{}, fmt.Errorf("failed to create project: %v", err)
	}

	return i, nil
}

// UpdateProject creates a GNS3 project with the specified name.
func UpdateProject(t Transport, id string, p ProjectUpdate) (ProjectInstance, error) {
	if id == "" {
		return ProjectInstance{}, errors.New("failed to update project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id)
	i := ProjectInstance{}
	_, err := t.Put(u, p, &i)
	if err != nil {
		return ProjectInstance{}, fmt.Errorf("failed to update project: %v", err)
	}

	return i, nil
}

// DeleteProject deletes a GNS3 project instance with the specified id.
func DeleteProject(t Transport, id string) (ProjectInstance, error) {
	if id == "" {
		return ProjectInstance{}, errors.New("failed to delete project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id)
	i := ProjectInstance{}
	_, err := t.Delete(u, &i)
	if err != nil {
		return ProjectInstance{}, fmt.Errorf("failed to delete project: %v", err)
	}

	return i, nil
}

// GetProject gets a GNS3 project instance with the specified id.
func GetProject(t Transport, id string) (ProjectInstance, error) {
	if id == "" {
		return ProjectInstance{}, errors.New("failed to get project: invalid project id")
	}

	u := "/v2/projects/" + url.PathEscape(id)
	i := ProjectInstance{}
	_, err := t.Get(u, &i)
	if err != nil {
		return ProjectInstance{}, fmt.Errorf("failed to get project: %v", err)
	}

	return i, nil
}

// GetProjects gets all the GNS3 projects.
func GetProjects(t Transport) ([]ProjectInstance, error) {
	u := "/v2/projects/"
	i := []ProjectInstance{}
	_, err := t.Get(u, &i)
	if err != nil {
		return []ProjectInstance{}, fmt.Errorf("failed to get projects: %v", err)
	}

	return i, nil
}
