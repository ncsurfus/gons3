package gns3tests

import (
	"gons3"
	"testing"
)

func TestUpdateProjectA(t *testing.T) {
	c := gons3.ProjectCreate{}
	c.SetName("TestUpdateA")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	defer gons3.DeleteProject(client, ci.ProjectID)

	u := gons3.ProjectUpdate{}
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
	c := gons3.ProjectCreate{}
	c.SetName("TestUpdateB")
	ci, err := gons3.CreateProject(client, c)
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	defer gons3.DeleteProject(client, ci.ProjectID)

	u := gons3.ProjectUpdate{}
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
