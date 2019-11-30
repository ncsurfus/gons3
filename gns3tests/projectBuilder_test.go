package gns3tests

import (
	"gons3"
	"testing"
)

func TestCreateProjectA(t *testing.T) {
	projectBuilder := gons3.ProjectBuilder{}
	projectBuilder.SetName("TestCreateA")
	projectBuilder.SetAutoClose(true)
	projectBuilder.SetSceneHeight(1800)
	projectBuilder.SetSceneWidth(900)
	projectBuilder.SetZoom(50)
	projectBuilder.SetShowLayers(true)
	projectBuilder.SetSnapToGrid(true)
	projectBuilder.SetShowGrid(true)
	projectBuilder.SetGridSize(10)
	projectBuilder.SetShowInterfaceLabels(true)
	projectBuilder.SetSupplier("testLogo", "https://example")
	projectBuilder.SetVariables([]gons3.ProjectVariables{
		gons3.ProjectVariables{
			Name:  "Name1",
			Value: "Value1",
		},
	})
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	if project.Name != "TestCreateA" {
		t.Errorf("Expected name: %v, got %v", "TestCreateA", project.Name)
	}
	if project.AutoClose != true {
		t.Errorf("Expected autoClose: %v, got %v", true, project.AutoClose)
	}
	if project.SceneHeight != 1800 {
		t.Errorf("Expected sceneHeight: %v, got %v", 1800, project.SceneHeight)
	}
	if project.SceneWidth != 900 {
		t.Errorf("Expected sceneWidth: %v, got %v", 900, project.SceneWidth)
	}
	if project.Zoom != 50 {
		t.Errorf("Expected zoom: %v, got %v", 50, project.Zoom)
	}
	if project.ShowLayers != true {
		t.Errorf("Expected showLayers: %v, got %v", true, project.ShowLayers)
	}
	if project.SnapToGrid != true {
		t.Errorf("Expected snapToGrid: %v, got %v", true, project.SnapToGrid)
	}
	if project.ShowGrid != true {
		t.Errorf("Expected showGrid: %v, got %v", true, project.ShowGrid)
	}
	if project.GridSize != 10 {
		t.Errorf("Expected gridSize: %v, got %v", 10, project.GridSize)
	}
	if project.ShowInterfaceLabels != true {
		t.Errorf("Expected showInterfaceLabels: %v, got %v", true, project.ShowInterfaceLabels)
	}
	if project.Supplier.Logo != "testLogo" {
		t.Errorf("Expected supplier logo: %v, got %v", "testLogo", project.Supplier.Logo)
	}
	if project.Supplier.URL != "https://example" {
		t.Errorf("Expected supplier URL: %v, got %v", "https://example", project.Supplier.URL)
	}
	if variables := *project.Variables; variables[0].Name != "Name1" {
		t.Errorf("Expected variables: %v, got %v", "Name1", variables[0].Name)
	}
	if variables := *project.Variables; variables[0].Value != "Value1" {
		t.Errorf("Expected variables: %v, got %v", "Value1", variables[0].Value)
	}
}

func TestCreateProjectB(t *testing.T) {
	projectBuilder := gons3.ProjectBuilder{}
	projectBuilder.SetName("TestCreateB")
	projectBuilder.SetAutoClose(false)
	projectBuilder.SetSceneHeight(1900)
	projectBuilder.SetSceneWidth(950)
	projectBuilder.SetZoom(100)
	projectBuilder.SetShowLayers(false)
	projectBuilder.SetSnapToGrid(false)
	projectBuilder.SetShowGrid(false)
	projectBuilder.SetGridSize(20)
	projectBuilder.SetShowInterfaceLabels(false)
	project, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("Error creating project: %v", err)
	}
	defer gons3.DeleteProject(client, project.ProjectID)

	if project.Name != "TestCreateB" {
		t.Errorf("Expected name: %v, got %v", "TestCreateB", project.Name)
	}
	if project.AutoClose != false {
		t.Errorf("Expected autoClose: %v, got %v", false, project.AutoClose)
	}
	if project.SceneHeight != 1900 {
		t.Errorf("Expected sceneHeight: %v, got %v", 1900, project.SceneHeight)
	}
	if project.SceneWidth != 950 {
		t.Errorf("Expected sceneWidth: %v, got %v", 950, project.SceneWidth)
	}
	if project.Zoom != 100 {
		t.Errorf("Expected zoom: %v, got %v", 100, project.Zoom)
	}
	if project.ShowLayers != false {
		t.Errorf("Expected showLayers: %v, got %v", false, project.ShowLayers)
	}
	if project.SnapToGrid != false {
		t.Errorf("Expected snapToGrid: %v, got %v", false, project.SnapToGrid)
	}
	if project.ShowGrid != false {
		t.Errorf("Expected showGrid: %v, got %v", false, project.ShowGrid)
	}
	if project.GridSize != 20 {
		t.Errorf("Expected gridSize: %v, got %v", 20, project.GridSize)
	}
	if project.ShowInterfaceLabels != false {
		t.Errorf("Expected showInterfaceLabels: %v, got %v", false, project.ShowInterfaceLabels)
	}
	if project.Supplier != nil {
		t.Errorf("Expected supplier: %v, got %v", nil, project.Supplier)
	}
	if project.Variables != nil {
		t.Errorf("Expected variables: %v, got %v", nil, project.Variables)
	}
}
