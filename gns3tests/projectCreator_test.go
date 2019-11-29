package gns3tests

import (
	"gons3"
	"testing"
)

func TestCreateProjectA(t *testing.T) {
	p := gons3.ProjectCreate{}
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
	defer gons3.DeleteProject(client, i.ProjectID)

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
	p := gons3.ProjectCreate{}
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
	defer gons3.DeleteProject(client, i.ProjectID)

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
