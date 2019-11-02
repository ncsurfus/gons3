// https://github.com/GNS3/gns3-server/blob/2.2/gns3server/schemas/project.py

package gons3

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
	Path                *string             `json:"path"`
	AutoClose           bool                `json:"auto_close"`
	ProjectID           *string             `json:"project_id"`
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

// UpdateProject models an update to a GNS3 project.
type UpdateProject struct {
	Name                *string             `json:"name"`
	Path                *string             `json:"path"`
	AutoClose           bool                `json:"auto_close"`
	AutoOpen            bool                `json:"auto_open"`
	AutoStart           bool                `json:"auto_start"`
	ProjectID           *string             `json:"project_id"`
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

// Project models an instance of a GNS3 project.
type Project struct {
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
