package gns3tests

import (
	"gons3"
	"testing"
)

var client = gons3.GNS3HTTPClient{}

func errorAssert(t *testing.T, name string, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v to be %v, got %v", name, expected, actual)
	}
}

func fatalAssert(t *testing.T, name string, expected, actual interface{}) {
	if expected != actual {
		t.Fatalf("Expected %v to be %v, got %v", name, expected, actual)
	}
}
