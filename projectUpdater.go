package gons3

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
