package gns3tests

import (
	"gons3"
	"testing"
)

func TestUpdateProjectA(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestUpdateA")
	newProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	defer gons3.DeleteProject(client, newProject.ProjectID)

	projectUpdater := gons3.ProjectUpdater{}
	projectUpdater.SetName("TestUpdateAA")
	projectUpdater.SetAutoClose(true)
	projectUpdater.SetAutoOpen(true)
	projectUpdater.SetAutoStart(true)
	projectUpdater.SetSceneHeight(1800)
	projectUpdater.SetSceneWidth(900)
	projectUpdater.SetZoom(50)
	projectUpdater.SetShowLayers(true)
	projectUpdater.SetSnapToGrid(true)
	projectUpdater.SetShowGrid(true)
	projectUpdater.SetGridSize(10)
	projectUpdater.SetShowInterfaceLabels(true)
	projectUpdater.SetSupplier("testLogo", "https://example")
	projectUpdater.SetVariables([]gons3.ProjectVariables{
		gons3.ProjectVariables{
			Name:  "Name1",
			Value: "Value1",
		},
	})
	project, err := gons3.UpdateProject(client, newProject.ProjectID, projectUpdater)
	if err != nil {
		t.Fatalf("failed to update project: %v", err)
	}

	if project.Name != "TestUpdateAA" {
		t.Errorf("Expected name: %v, got %v", "TestUpdateAA", project.Name)
	}
	if project.AutoClose != true {
		t.Errorf("Expected autoClose: %v, got %v", true, project.AutoClose)
	}
	if project.AutoOpen != true {
		t.Errorf("Expected autoOpen: %v, got %v", true, project.AutoClose)
	}
	if project.AutoStart != true {
		t.Errorf("Expected autoStart: %v, got %v", true, project.AutoClose)
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

func TestUpdateProjectB(t *testing.T) {
	projectBuilder := gons3.NewProjectBuilder("TestUpdateB")
	newProject, err := gons3.CreateProject(client, projectBuilder)
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	defer gons3.DeleteProject(client, newProject.ProjectID)

	projectUpdater := gons3.ProjectUpdater{}
	projectUpdater.SetName("TestUpdateBB")
	projectUpdater.SetAutoClose(false)
	projectUpdater.SetAutoOpen(false)
	projectUpdater.SetAutoStart(false)
	projectUpdater.SetSceneHeight(1750)
	projectUpdater.SetSceneWidth(850)
	projectUpdater.SetZoom(100)
	projectUpdater.SetShowLayers(false)
	projectUpdater.SetSnapToGrid(false)
	projectUpdater.SetShowGrid(false)
	projectUpdater.SetGridSize(20)
	projectUpdater.SetShowInterfaceLabels(false)
	project, err := gons3.UpdateProject(client, newProject.ProjectID, projectUpdater)
	if err != nil {
		t.Fatalf("failed to update project: %v", err)
	}

	if project.Name != "TestUpdateBB" {
		t.Errorf("Expected name: %v, got %v", "TestUpdateAA", project.Name)
	}
	if project.AutoClose != false {
		t.Errorf("Expected autoClose: %v, got %v", false, project.AutoClose)
	}
	if project.AutoOpen != false {
		t.Errorf("Expected autoOpen: %v, got %v", false, project.AutoClose)
	}
	if project.AutoStart != false {
		t.Errorf("Expected autoStart: %v, got %v", false, project.AutoClose)
	}
	if project.SceneHeight != 1750 {
		t.Errorf("Expected sceneHeight: %v, got %v", 1750, project.SceneHeight)
	}
	if project.SceneWidth != 850 {
		t.Errorf("Expected sceneWidth: %v, got %v", 850, project.SceneWidth)
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
		t.Errorf("Expected varaibles: %v, got %v", nil, project.Variables)
	}
}
