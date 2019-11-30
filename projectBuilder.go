package gons3

// ProjectBuilder models a new GNS3 project.
type ProjectBuilder struct {
	values map[string]interface{}
}

// NewProjectBuilder initializes a new project with the specified name.
func NewProjectBuilder(name string) ProjectBuilder {
	projectBuilder := ProjectBuilder{}
	projectBuilder.SetName(name)
	return projectBuilder
}

// SetProperty sets a custom property and value for the project.
func (n *ProjectBuilder) SetProperty(name string, value interface{}) {
	if n.values == nil {
		n.values = map[string]interface{}{}
	}
	n.values[name] = value
}

// SetName sets the name for the new project.
func (n *ProjectBuilder) SetName(name string) {
	n.SetProperty("name", name)
}

// SetPath sets the path for the new project.
func (n *ProjectBuilder) SetPath(path string) {
	n.SetProperty("path", path)
}

// SetAutoClose sets the auto_close option for the new project.
func (n *ProjectBuilder) SetAutoClose(autoClose bool) {
	n.SetProperty("auto_close", autoClose)
}

// SetProjectID sets the project_id option for the new project.
func (n *ProjectBuilder) SetProjectID(projectID string) {
	n.SetProperty("project_id", projectID)
}

// SetSceneHeight sets the scene_height option for the new project.
func (n *ProjectBuilder) SetSceneHeight(height int) {
	n.SetProperty("scene_height", height)
}

// SetSceneWidth sets the scene_width option for the new project.
func (n *ProjectBuilder) SetSceneWidth(width int) {
	n.SetProperty("scene_width", width)
}

// SetZoom sets the zoom option for the new project.
func (n *ProjectBuilder) SetZoom(zoom int) {
	n.SetProperty("zoom", zoom)
}

// SetShowLayers sets the show_layers option for the new project.
func (n *ProjectBuilder) SetShowLayers(showLayers bool) {
	n.SetProperty("show_layers", showLayers)
}

// SetSnapToGrid sets the snap_to_grid option for the new project.
func (n *ProjectBuilder) SetSnapToGrid(snapToGrid bool) {
	n.SetProperty("snap_to_grid", snapToGrid)
}

// SetShowGrid sets the show_grid option for the new project.
func (n *ProjectBuilder) SetShowGrid(showGrid bool) {
	n.SetProperty("show_grid", showGrid)
}

// SetGridSize sets the grid_size option for the new project.
func (n *ProjectBuilder) SetGridSize(gridSize int) {
	n.SetProperty("grid_size", gridSize)
}

// SetShowInterfaceLabels sets the show_interface_labels option for the new project.
func (n *ProjectBuilder) SetShowInterfaceLabels(showInterfaceLabels bool) {
	n.SetProperty("show_interface_labels", showInterfaceLabels)
}

// SetSupplier sets the supplier option for the new project.
func (n *ProjectBuilder) SetSupplier(logo string, url string) {
	supplier := map[string]interface{}{
		"logo": logo,
		"url":  url,
	}
	n.SetProperty("supplier", supplier)
}

// SetVariables sets the variables option for the new project.
func (n *ProjectBuilder) SetVariables(variables []ProjectVariables) {
	n.SetProperty("variables", variables)
}
