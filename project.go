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
		return Project{}, ErrEmptyProjectID
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
		return ErrEmptyProjectID
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
		return Project{}, ErrEmptyProjectID
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
		return Project{}, ErrEmptyProjectID
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
		return Project{}, ErrEmptyProjectID
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
		return []byte{}, ErrEmptyProjectID
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
		return ErrEmptyProjectID
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

// ProjectCreator models a new GNS3 project.
type ProjectCreator struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the project.
func (n *ProjectCreator) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetName sets the name for the new project.
func (n *ProjectCreator) SetName(name string) {
	n.SetProperty("name", name)
}

// SetPath sets the path for the new project.
func (n *ProjectCreator) SetPath(path string) {
	n.SetProperty("path", path)
}

// SetAutoClose sets the auto_close option for the new project.
func (n *ProjectCreator) SetAutoClose(autoClose bool) {
	n.SetProperty("auto_close", autoClose)
}

// SetProjectID sets the project_id option for the new project.
func (n *ProjectCreator) SetProjectID(projectID string) {
	n.SetProperty("project_id", projectID)
}

// SetSceneHeight sets the scene_height option for the new project.
func (n *ProjectCreator) SetSceneHeight(height int) {
	n.SetProperty("scene_height", height)
}

// SetSceneWidth sets the scene_width option for the new project.
func (n *ProjectCreator) SetSceneWidth(width int) {
	n.SetProperty("scene_width", width)
}

// SetZoom sets the zoom option for the new project.
func (n *ProjectCreator) SetZoom(zoom int) {
	n.SetProperty("zoom", zoom)
}

// SetShowLayers sets the show_layers option for the new project.
func (n *ProjectCreator) SetShowLayers(showLayers bool) {
	n.SetProperty("show_layers", showLayers)
}

// SetSnapToGrid sets the snap_to_grid option for the new project.
func (n *ProjectCreator) SetSnapToGrid(snapToGrid bool) {
	n.SetProperty("snap_to_grid", snapToGrid)
}

// SetShowGrid sets the show_grid option for the new project.
func (n *ProjectCreator) SetShowGrid(showGrid bool) {
	n.SetProperty("show_grid", showGrid)
}

// SetGridSize sets the grid_size option for the new project.
func (n *ProjectCreator) SetGridSize(gridSize int) {
	n.SetProperty("grid_size", gridSize)
}

// SetShowInterfaceLabels sets the show_interface_labels option for the new project.
func (n *ProjectCreator) SetShowInterfaceLabels(showInterfaceLabels bool) {
	n.SetProperty("show_interface_labels", showInterfaceLabels)
}

// SetSupplier sets the supplier option for the new project.
func (n *ProjectCreator) SetSupplier(logo string, url string) {
	supplier := map[string]interface{}{
		"logo": logo,
		"url":  url,
	}
	n.SetProperty("supplier", supplier)
}

// SetVariables sets the variables option for the new project.
func (n *ProjectCreator) SetVariables(variables []ProjectVariables) {
	n.SetProperty("variables", variables)
}

// ProjectUpdater models an update to a GNS3 project.
type ProjectUpdater struct {
	values map[string]interface{}
}

// SetProperty sets a custom property and value for the project.
func (n *ProjectUpdater) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetName sets the name for the project.
func (n *ProjectUpdater) SetName(name string) {
	n.SetProperty("name", name)
}

// SetPath sets the path for the project.
func (n *ProjectUpdater) SetPath(path string) {
	n.SetProperty("path", path)
}

// SetAutoClose sets the auto_close option for the project.
func (n *ProjectUpdater) SetAutoClose(autoClose bool) {
	n.SetProperty("auto_close", autoClose)
}

// SetAutoOpen sets the auto_open option for the project.
func (n *ProjectUpdater) SetAutoOpen(autoOpen bool) {
	n.SetProperty("auto_open", autoOpen)
}

// SetAutoStart sets the auto_start option for the project.
func (n *ProjectUpdater) SetAutoStart(autoStart bool) {
	n.SetProperty("auto_start", autoStart)
}

// SetSceneHeight sets the scene_height option for the project.
func (n *ProjectUpdater) SetSceneHeight(height int) {
	n.SetProperty("scene_height", height)
}

// SetSceneWidth sets the scene_width option for the project.
func (n *ProjectUpdater) SetSceneWidth(width int) {
	n.SetProperty("scene_width", width)
}

// SetZoom sets the zoom option for the project.
func (n *ProjectUpdater) SetZoom(zoom int) {
	n.SetProperty("zoom", zoom)
}

// SetShowLayers sets the show_layers option for the project.
func (n *ProjectUpdater) SetShowLayers(showLayers bool) {
	n.SetProperty("show_layers", showLayers)
}

// SetSnapToGrid sets the snap_to_grid option for the project.
func (n *ProjectUpdater) SetSnapToGrid(snapToGrid bool) {
	n.SetProperty("snap_to_grid", snapToGrid)
}

// SetShowGrid sets the show_grid option for the project.
func (n *ProjectUpdater) SetShowGrid(showGrid bool) {
	n.SetProperty("show_grid", showGrid)
}

// SetGridSize sets the grid_size option for the project.
func (n *ProjectUpdater) SetGridSize(gridSize int) {
	n.SetProperty("grid_size", gridSize)
}

// SetShowInterfaceLabels sets the show_interface_labels option for the project.
func (n *ProjectUpdater) SetShowInterfaceLabels(showInterfaceLabels bool) {
	n.SetProperty("show_interface_labels", showInterfaceLabels)
}

// SetSupplier sets the supplier option for the project.
func (n *ProjectUpdater) SetSupplier(logo string, url string) {
	supplier := map[string]interface{}{
		"logo": logo,
		"url":  url,
	}
	n.SetProperty("supplier", supplier)
}

// RemoveSupplier clears the supplier option for the project.
func (n *ProjectUpdater) RemoveSupplier() {
	n.SetProperty("supplier", nil)
}

// SetVariables sets the variables option for the project.
func (n *ProjectUpdater) SetVariables(variables []ProjectVariables) {
	n.SetProperty("variables", variables)
}

// RemoveVariables clears the variables option for the project.
func (n *ProjectUpdater) RemoveVariables() {
	n.SetProperty("variables", nil)
}
