package expectations

import (
	"gons3"
	"testing"
)

func TestUpdateProject_A(t *testing.T) {
	g := gons3.GNS3HTTPClient{}

	err := deleteProjectByName(g, "TestUpdateA")
	if err != nil {
		t.Fatal(err)
	}

	c := gons3.ProjectCreator{}
	c.SetName("TestUpdateA")
	ci, err := gons3.CreateProject(g, c)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteProjectByName(g, "TestUpdateAA")

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
	i, err := gons3.UpdateProject(g, ci.ProjectID, u)
	if err != nil {
		t.Fatal(err)
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
