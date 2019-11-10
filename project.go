// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/schemas/project.py

package gons3

import (
	"errors"
	"fmt"
	"net/url"
)

const jsonEncoding = "applicaton/json"

// ErrInvalidResponse means there was an error reading or parsing the http resonse.
var ErrInvalidResponse = errors.New("the response was could not be read or parsed")

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
func CreateProject(g GNS3Client, p ProjectCreator) (Project, error) {
	path := "/v2/projects"
	proj := Project{}
	if err := post(g, path, 201, p.values, &proj); err != nil {
		return Project{}, fmt.Errorf("http request failed: %w", err)
	}
	return proj, nil
}

// UpdateProject creates a GNS3 project with the specified name.
func UpdateProject(g GNS3Client, id string, p ProjectUpdater) (Project, error) {
	if id == "" {
		return Project{}, errors.New("failed to update project: invalid project id")
	}

	path := "/v2/projects/" + url.PathEscape(id)
	proj := Project{}
	if err := put(g, path, 200, p.values, &proj); err != nil {
		return Project{}, fmt.Errorf("failed to update project: %v", err)
	}
	return proj, nil
}

// DeleteProject deletes a GNS3 project instance with the specified id.
func DeleteProject(g GNS3Client, id string) error {
	if id == "" {
		return errors.New("failed to delete project: invalid project id")
	}

	path := "/v2/projects/" + url.PathEscape(id)
	if err := delete(g, path, 204, nil); err != nil {
		return fmt.Errorf("failed to delete project: %v", err)
	}
	return nil
}

// GetProject gets a GNS3 project instance with the specified id.
func GetProject(g GNS3Client, id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("failed to get project: invalid project id")
	}

	path := "/v2/projects/" + url.PathEscape(id)
	proj := Project{}
	if err := get(g, path, 200, &proj); err != nil {
		return Project{}, fmt.Errorf("failed to get project: %v", err)
	}
	return proj, nil
}

// GetProjects gets all the GNS3 projects.
func GetProjects(g GNS3Client) ([]Project, error) {
	path := "/v2/projects"
	proj := []Project{}
	if err := get(g, path, 200, &proj); err != nil {
		return []Project{}, fmt.Errorf("failed to get projects: %v", err)
	}
	return proj, nil
}

// OpenProject opens the GNS3 project.
func OpenProject(g GNS3Client, id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("failed to open project: invalid project id")
	}

	path := "/v2/projects/" + url.PathEscape(id) + "/open"
	proj := Project{}
	if err := post(g, path, 201, nil, &proj); err != nil {
		return Project{}, fmt.Errorf("failed to open project: %v", err)
	}
	return proj, nil
}

// CloseProject opens the GNS3 project.
func CloseProject(g GNS3Client, id string) (Project, error) {
	if id == "" {
		return Project{}, errors.New("failed to close project: invalid project id")
	}

	path := "/v2/projects/" + url.PathEscape(id) + "/close"
	proj := Project{}
	if err := post(g, path, 204, nil, &proj); err != nil {
		return Project{}, fmt.Errorf("failed to close project: %v", err)
	}
	return proj, nil
}

// ReadProjectFile reads a GNS3 project's file.
func ReadProjectFile(g GNS3Client, id string, filepath string) ([]byte, error) {
	if id == "" {
		return []byte{}, errors.New("failed to read project file: invalid project id")
	}
	if filepath == "" {
		return []byte{}, errors.New("failed to read project file: invalid path")
	}

	path := "/v2/projects/" + url.PathEscape(id) + "/files/" + filepath
	data := []byte{}
	if err := get(g, path, 200, &data); err != nil {
		return []byte{}, fmt.Errorf("failed to close project: %v", err)
	}
	return data, nil
}

// WriteProjectFile writes a GNS3 project's file.
func WriteProjectFile(g GNS3Client, id string, filepath string, data []byte) error {
	if id == "" {
		return errors.New("failed to write project file: invalid project id")
	}
	if filepath == "" {
		return errors.New("failed to write project file: invalid path")
	}

	path := "/v2/projects/" + url.PathEscape(id) + "/files/" + filepath
	if err := post(g, path, 200, &data, nil); err != nil {
		return fmt.Errorf("failed to close project: %v", err)
	}
	return nil
}
