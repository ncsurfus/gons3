// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/schemas/project.py
// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/handlers/api/controller/project_handler.py

package gons3

import (
	"errors"
	"net/url"
)

// ErrEmptyProjectID means that the project id cannot be empty.
var ErrEmptyProjectID = errors.New("projectID cannot be empty")

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
func (project Project) IsOpened() bool {
	return project.Status == "opened"
}

// CreateProject creates a GNS3 project with the specified name.
func CreateProject(client GNS3Client, projectBuilder ProjectBuilder) (Project, error) {
	path := "/v2/projects"
	project := Project{}
	if err := post(client, path, 201, projectBuilder.values, &project); err != nil {
		return Project{}, err
	}
	return project, nil
}

// UpdateProject creates a GNS3 project with the specified name.
func UpdateProject(client GNS3Client, projectID string, projectUpdater ProjectUpdater) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID)
	project := Project{}
	if err := put(client, path, 200, projectUpdater.values, &project); err != nil {
		return Project{}, err
	}
	return project, nil
}

// DeleteProject deletes a GNS3 project instance with the specified id.
func DeleteProject(client GNS3Client, projectID string) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID)
	if err := delete(client, path, 204, nil); err != nil {
		return err
	}
	return nil
}

// GetProject gets a GNS3 project instance with the specified id.
func GetProject(client GNS3Client, projectID string) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID)
	project := Project{}
	if err := get(client, path, 200, &project); err != nil {
		return Project{}, err
	}
	return project, nil
}

// GetProjects gets all the GNS3 projects.
func GetProjects(client GNS3Client) ([]Project, error) {
	path := "/v2/projects"
	project := []Project{}
	if err := get(client, path, 200, &project); err != nil {
		return []Project{}, err
	}
	return project, nil
}

// OpenProject opens the GNS3 project.
func OpenProject(client GNS3Client, projectID string) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/open"
	project := Project{}
	if err := post(client, path, 201, nil, &project); err != nil {
		return Project{}, err
	}
	return project, nil
}

// CloseProject opens the GNS3 project.
func CloseProject(client GNS3Client, projectID string) (Project, error) {
	if projectID == "" {
		return Project{}, ErrEmptyProjectID
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/close"
	PROJECT := Project{}
	if err := post(client, path, 201, nil, &PROJECT); err != nil {
		return Project{}, err
	}
	return PROJECT, nil
}

// ReadProjectFile reads a GNS3 project's file.
func ReadProjectFile(client GNS3Client, projectID, filepath string) ([]byte, error) {
	if projectID == "" {
		return []byte{}, ErrEmptyProjectID
	}
	if filepath == "" {
		return []byte{}, ErrEmptyFilepath
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/files/" + filepath
	data := []byte{}
	if err := get(client, path, 200, &data); err != nil {
		return []byte{}, err
	}
	return data, nil
}

// WriteProjectFile writes a GNS3 project's file.
func WriteProjectFile(client GNS3Client, projectID, filepath string, data []byte) error {
	if projectID == "" {
		return ErrEmptyProjectID
	}
	if filepath == "" {
		return ErrEmptyFilepath
	}

	path := "/v2/projects/" + url.PathEscape(projectID) + "/files/" + filepath
	if err := post(client, path, 200, &data, nil); err != nil {
		return err
	}
	return nil
}
