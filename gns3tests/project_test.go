package gns3tests

import (
	"gons3"
	"testing"
)

func TestGetProject(t *testing.T) {
	c := gons3.ProjectCreator{}
	c.SetName("TestGetProject")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, ci.ProjectID)

	proj, err := gons3.GetProject(client, ci.ProjectID)
	if err != nil {
		t.Fatalf("Error getting project: %v", err)
	}
	if proj.Name != "TestGetProject" {
		t.Errorf("Expected name: %v, got %v", "TestGetProject", proj.Name)
	}
}

func TestOpenCloseProject(t *testing.T) {
	c := gons3.ProjectCreator{}
	c.SetName("TestOpenCloseProject")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, ci.ProjectID)

	ci, err = gons3.OpenProject(client, ci.ProjectID)
	if err != nil {
		t.Fatalf("Error opening project: %v", err)
	}
	if !ci.IsOpened() {
		t.Errorf("Expected IsOpened(): %v, got %v", true, ci.IsOpened())
	}

	ci, err = gons3.CloseProject(client, ci.ProjectID)
	if err != nil {
		t.Fatalf("Error closing project: %v", err)
	}
	if ci.IsOpened() {
		t.Errorf("Expected IsOpened(): %v, got %v", false, ci.IsOpened())
	}
}

func TestReadWriteProjectFile(t *testing.T) {
	c := gons3.ProjectCreator{}
	c.SetName("TestReadWriteProjectFile")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, ci.ProjectID)

	err = gons3.WriteProjectFile(client, ci.ProjectID, "testing", []byte("the test"))
	if err != nil {
		t.Fatalf("Error writing project file: %v", err)
	}

	data, err := gons3.ReadProjectFile(client, ci.ProjectID, "testing")
	sdata := string(data)
	if err != nil {
		t.Fatalf("Error reading project file: %v", err)
	}
	if sdata != "the test" {
		t.Errorf("Expected data: %v, got %v", "the test", sdata)
	}
}

func TestDeleteProject(t *testing.T) {
	c := gons3.ProjectCreator{}
	c.SetName("TestDeleteProject")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}

	err = gons3.DeleteProject(client, ci.ProjectID)
	if err != nil {
		t.Fatalf("Error deleting project: %v", err)
	}
}


func TestCreateProjectA(t *testing.T) {
	p := gons3.ProjectCreator{}
	p.SetName("TestCreateA")
	p.SetAutoClose(true)
	p.SetSceneHeight(1800)
	p.SetSceneWidth(900)
	p.SetZoom(50)
	p.SetShowLayers(true)
	p.SetSnapToGrid(true)
	p.SetShowGrid(true)
	p.SetGridSize(10)
	p.SetShowInterfaceLabels(true)
	p.SetSupplier("testLogo", "https://example")
	p.SetVariables([]gons3.ProjectVariables{
		gons3.ProjectVariables{
			Name:  "Name1",
			Value: "Value1",
		},
	})
	i, err := gons3.CreateProject(client, p)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer deleteProjectByName(client, "TestCreateA")

	if i.Name != "TestCreateA" {
		t.Errorf("Expected name: %v, got %v", "TestCreateA", i.Name)
	}
	if i.AutoClose != true {
		t.Errorf("Expected autoClose: %v, got %v", true, i.AutoClose)
	}
	if i.SceneHeight != 1800 {
		t.Errorf("Expected sceneHeight: %v, got %v", 1800, i.SceneHeight)
	}
	if i.SceneWidth != 900 {
		t.Errorf("Expected sceneWidth: %v, got %v", 900, i.SceneWidth)
	}
	if i.Zoom != 50 {
		t.Errorf("Expected zoom: %v, got %v", 50, i.Zoom)
	}
	if i.ShowLayers != true {
		t.Errorf("Expected showLayers: %v, got %v", true, i.ShowLayers)
	}
	if i.SnapToGrid != true {
		t.Errorf("Expected snapToGrid: %v, got %v", true, i.SnapToGrid)
	}
	if i.ShowGrid != true {
		t.Errorf("Expected showGrid: %v, got %v", true, i.ShowGrid)
	}
	if i.GridSize != 10 {
		t.Errorf("Expected gridSize: %v, got %v", 10, i.GridSize)
	}
	if i.ShowInterfaceLabels != true {
		t.Errorf("Expected showInterfaceLabels: %v, got %v", true, i.ShowInterfaceLabels)
	}
	if i.Supplier.Logo != "testLogo" {
		t.Errorf("Expected supplier logo: %v, got %v", "testLogo", i.Supplier.Logo)
	}
	if i.Supplier.URL != "https://example" {
		t.Errorf("Expected supplier URL: %v, got %v", "https://example", i.Supplier.URL)
	}
	if variables := *i.Variables; variables[0].Name != "Name1" {
		t.Errorf("Expected variables: %v, got %v", "Name1", variables[0].Name)
	}
	if variables := *i.Variables; variables[0].Value != "Value1" {
		t.Errorf("Expected variables: %v, got %v", "Value1", variables[0].Value)
	}
}

func TestCreateProjectB(t *testing.T) {
	p := gons3.ProjectCreator{}
	p.SetName("TestCreateB")
	p.SetAutoClose(false)
	p.SetSceneHeight(1900)
	p.SetSceneWidth(950)
	p.SetZoom(100)
	p.SetShowLayers(false)
	p.SetSnapToGrid(false)
	p.SetShowGrid(false)
	p.SetGridSize(20)
	p.SetShowInterfaceLabels(false)
	i, err := gons3.CreateProject(client, p)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer deleteProjectByName(client, "TestCreateB")

	if i.Name != "TestCreateB" {
		t.Errorf("Expected name: %v, got %v", "TestCreateB", i.Name)
	}
	if i.AutoClose != false {
		t.Errorf("Expected autoClose: %v, got %v", false, i.AutoClose)
	}
	if i.SceneHeight != 1900 {
		t.Errorf("Expected sceneHeight: %v, got %v", 1900, i.SceneHeight)
	}
	if i.SceneWidth != 950 {
		t.Errorf("Expected sceneWidth: %v, got %v", 950, i.SceneWidth)
	}
	if i.Zoom != 100 {
		t.Errorf("Expected zoom: %v, got %v", 100, i.Zoom)
	}
	if i.ShowLayers != false {
		t.Errorf("Expected showLayers: %v, got %v", false, i.ShowLayers)
	}
	if i.SnapToGrid != false {
		t.Errorf("Expected snapToGrid: %v, got %v", false, i.SnapToGrid)
	}
	if i.ShowGrid != false {
		t.Errorf("Expected showGrid: %v, got %v", false, i.ShowGrid)
	}
	if i.GridSize != 20 {
		t.Errorf("Expected gridSize: %v, got %v", 20, i.GridSize)
	}
	if i.ShowInterfaceLabels != false {
		t.Errorf("Expected showInterfaceLabels: %v, got %v", false, i.ShowInterfaceLabels)
	}
	if i.Supplier != nil {
		t.Errorf("Expected supplier: %v, got %v", nil, i.Supplier)
	}
	if i.Variables != nil {
		t.Errorf("Expected variables: %v, got %v", nil, i.Variables)
	}
}


func TestUpdateProjectA(t *testing.T) {
	c := gons3.ProjectCreator{}
	c.SetName("TestUpdateA")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	defer deleteProjectByName(client, "TestUpdateAA")

	u := gons3.ProjectUpdater{}
	u.SetName("TestUpdateAA")
	u.SetAutoClose(true)
	u.SetAutoOpen(true)
	u.SetAutoStart(true)
	u.SetSceneHeight(1800)
	u.SetSceneWidth(900)
	u.SetZoom(50)
	u.SetShowLayers(true)
	u.SetSnapToGrid(true)
	u.SetShowGrid(true)
	u.SetGridSize(10)
	u.SetShowInterfaceLabels(true)
	u.SetSupplier("testLogo", "https://example")
	u.SetVariables([]gons3.ProjectVariables{
		gons3.ProjectVariables{
			Name:  "Name1",
			Value: "Value1",
		},
	})
	i, err := gons3.UpdateProject(client, ci.ProjectID, u)
	if err != nil {
		t.Fatalf("failed to update project: %v", err)
	}

	if i.Name != "TestUpdateAA" {
		t.Errorf("Expected name: %v, got %v", "TestUpdateAA", i.Name)
	}
	if i.AutoClose != true {
		t.Errorf("Expected autoClose: %v, got %v", true, i.AutoClose)
	}
	if i.AutoOpen != true {
		t.Errorf("Expected autoOpen: %v, got %v", true, i.AutoClose)
	}
	if i.AutoStart != true {
		t.Errorf("Expected autoStart: %v, got %v", true, i.AutoClose)
	}
	if i.SceneHeight != 1800 {
		t.Errorf("Expected sceneHeight: %v, got %v", 1800, i.SceneHeight)
	}
	if i.SceneWidth != 900 {
		t.Errorf("Expected sceneWidth: %v, got %v", 900, i.SceneWidth)
	}
	if i.Zoom != 50 {
		t.Errorf("Expected zoom: %v, got %v", 50, i.Zoom)
	}
	if i.ShowLayers != true {
		t.Errorf("Expected showLayers: %v, got %v", true, i.ShowLayers)
	}
	if i.SnapToGrid != true {
		t.Errorf("Expected snapToGrid: %v, got %v", true, i.SnapToGrid)
	}
	if i.ShowGrid != true {
		t.Errorf("Expected showGrid: %v, got %v", true, i.ShowGrid)
	}
	if i.GridSize != 10 {
		t.Errorf("Expected gridSize: %v, got %v", 10, i.GridSize)
	}
	if i.ShowInterfaceLabels != true {
		t.Errorf("Expected showInterfaceLabels: %v, got %v", true, i.ShowInterfaceLabels)
	}
	if i.Supplier.Logo != "testLogo" {
		t.Errorf("Expected supplier logo: %v, got %v", "testLogo", i.Supplier.Logo)
	}
	if i.Supplier.URL != "https://example" {
		t.Errorf("Expected supplier URL: %v, got %v", "https://example", i.Supplier.URL)
	}
	if variables := *i.Variables; variables[0].Name != "Name1" {
		t.Errorf("Expected variables: %v, got %v", "Name1", variables[0].Name)
	}
	if variables := *i.Variables; variables[0].Value != "Value1" {
		t.Errorf("Expected variables: %v, got %v", "Value1", variables[0].Value)
	}
}

func TestUpdateProjectB(t *testing.T) {
	c := gons3.ProjectCreator{}
	c.SetName("TestUpdateB")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	defer deleteProjectByName(client, "TestUpdateBB")

	u := gons3.ProjectUpdater{}
	u.SetName("TestUpdateBB")
	u.SetAutoClose(false)
	u.SetAutoOpen(false)
	u.SetAutoStart(false)
	u.SetSceneHeight(1750)
	u.SetSceneWidth(850)
	u.SetZoom(100)
	u.SetShowLayers(false)
	u.SetSnapToGrid(false)
	u.SetShowGrid(false)
	u.SetGridSize(20)
	u.SetShowInterfaceLabels(false)
	i, err := gons3.UpdateProject(client, ci.ProjectID, u)
	if err != nil {
		t.Fatalf("failed to update project: %v", err)
	}

	if i.Name != "TestUpdateBB" {
		t.Errorf("Expected name: %v, got %v", "TestUpdateAA", i.Name)
	}
	if i.AutoClose != false {
		t.Errorf("Expected autoClose: %v, got %v", false, i.AutoClose)
	}
	if i.AutoOpen != false {
		t.Errorf("Expected autoOpen: %v, got %v", false, i.AutoClose)
	}
	if i.AutoStart != false {
		t.Errorf("Expected autoStart: %v, got %v", false, i.AutoClose)
	}
	if i.SceneHeight != 1750 {
		t.Errorf("Expected sceneHeight: %v, got %v", 1750, i.SceneHeight)
	}
	if i.SceneWidth != 850 {
		t.Errorf("Expected sceneWidth: %v, got %v", 850, i.SceneWidth)
	}
	if i.Zoom != 100 {
		t.Errorf("Expected zoom: %v, got %v", 100, i.Zoom)
	}
	if i.ShowLayers != false {
		t.Errorf("Expected showLayers: %v, got %v", false, i.ShowLayers)
	}
	if i.SnapToGrid != false {
		t.Errorf("Expected snapToGrid: %v, got %v", false, i.SnapToGrid)
	}
	if i.ShowGrid != false {
		t.Errorf("Expected showGrid: %v, got %v", false, i.ShowGrid)
	}
	if i.GridSize != 20 {
		t.Errorf("Expected gridSize: %v, got %v", 20, i.GridSize)
	}
	if i.ShowInterfaceLabels != false {
		t.Errorf("Expected showInterfaceLabels: %v, got %v", false, i.ShowInterfaceLabels)
	}
	if i.Supplier != nil {
		t.Errorf("Expected supplier: %v, got %v", nil, i.Supplier)
	}
	if i.Variables != nil {
		t.Errorf("Expected varaibles: %v, got %v", nil, i.Variables)
	}
}
