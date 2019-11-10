// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/schemas/project.py
// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/handlers/api/controller/project_handler.py

package gons3

import (
	"errors"
	"net/url"
)

// ErrEmptyID means that the project id cannot be empty.
var ErrEmptyID = errors.New("id cannot be empty")

// ErrEmptyFilepath means that the filepath cannot be empty.
var ErrEmptyFilepath = errors.New("filepath cannot be empty")

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

// IsOpened returns true if the project status is set to opened.
func (proj Project) IsOpened() bool {
	return proj.Status == "opened"
}

// CreateProject creates a GNS3 project with the specified name.
func CreateProject(g GNS3Client, p ProjectCreator) (Project, error) {
	path := "/v2/projects"
	proj := Project{}
	if err := post(g, path, 201, p.values, &proj); err != nil {
		return Project{}, err
	}
	return proj, nil
}

// UpdateProject creates a GNS3 project with the specified name.
func UpdateProject(g GNS3Client, projectID string, p ProjectUpdater) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyID
	}

	path := "/v2/projects/" + url.PathEscape(projectID)
	proj := Project{}
	if err := put(g, path, 200, p.values, &proj); err != nil {
		return Project{}, err
	}
	return proj, nil
}

// DeleteProject deletes a GNS3 project instance with the specified id.
func DeleteProject(g GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyID
	}

	path := "/v2/projects/" + url.PathEscape(projectID)
	if err := delete(g, path, 204, nil); err != nil {
		return err
	}
	return nil
}

// GetProject gets a GNS3 project instance with the specified id.
func GetProject(g GNS3Client, projectID string) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyID
	}

	path := "/v2/projects/" + url.PathEscape(projectID)
	proj := Project{}
	if err := get(g, path, 200, &proj); err != nil {
		return Project{}, err
	}
	return proj, nil
}

// GetProjects gets all the GNS3 projects.
func GetProjects(g GNS3Client) ([]Project, error) {
	path := "/v2/projects"
	proj := []Project{}
	if err := get(g, path, 200, &proj); err != nil {
		return []Project{}, err
	}
	return proj, nil
}

// OpenProject opens the GNS3 project.
func OpenProject(g GNS3Client, projectID string) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/open"
	proj := Project{}
	if err := post(g, path, 201, nil, &proj); err != nil {
		return Project{}, err
	}
	return proj, nil
}

// CloseProject opens the GNS3 project.
func CloseProject(g GNS3Client, projectID string) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/close"
	proj := Project{}
	if err := post(g, path, 201, nil, &proj); err != nil {
		return Project{}, err
	}
	return proj, nil
}

// ReadProjectFile reads a GNS3 project's file.
func ReadProjectFile(g GNS3Client, projectID string, filepath string) ([]byte, error) {
	if projectID == "" {
		return []byte{}, ErrEmptyID
	}
	if filepath == "" {
		return []byte{}, ErrEmptyFilepath
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/files/" + filepath
	data := []byte{}
	if err := get(g, path, 200, &data); err != nil {
		return []byte{}, err
	}
	return data, nil
}

// WriteProjectFile writes a GNS3 project's file.
func WriteProjectFile(g GNS3Client, projectID string, filepath string, data []byte) error {
	if projectID == "" {
		return ErrEmptyID
	}
	if filepath == "" {
		return ErrEmptyFilepath
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/files/" + filepath
	if err := post(g, path, 200, &data, nil); err != nil {
		return err
	}
	return nil
}
