package loading_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/WrenchRobotics/urdf-go/loaders"
)

/*
TestFromURDFFile_missing_file_returns_does_not_exist_error
Description:

	Tests that FromURDFFile returns an error containing "does not exist"
	when the provided file path does not exist.
*/
func TestFromURDFFile_missing_file_returns_does_not_exist_error(t *testing.T) {
	// Setup
	nonExistentPath := filepath.Join(t.TempDir(), "missing.urdf")

	// Execute
	model, err := loaders.FromURDFFile(nonExistentPath)

	// Verify
	if err == nil {
		t.Fatalf("expected an error for a non-existent file path, got nil")
	}
	if model != nil {
		t.Fatalf("expected model to be nil when loading a non-existent file, got %#v", model)
	}
	if !strings.Contains(err.Error(), "does not exist") {
		t.Fatalf("expected error to contain 'does not exist', got %q", err.Error())
	}
}

/*
TestFromURDFFile_loads_example_urdf_and_detects_link_count
Description:

	Tests that FromURDFFile successfully loads a URDF from the examples
	directory and detects the expected number of links.
*/
func TestFromURDFFile_loads_example_urdf_and_detects_link_count(t *testing.T) {
	// Setup
	exampleURDFPath := filepath.Join("..", "..", "examples", "beaker", "500ml.urdf")

	// Execute
	model, err := loaders.FromURDFFile(exampleURDFPath)

	// Verify
	if err != nil {
		t.Fatalf("expected example URDF to load successfully, got error: %v", err)
	}
	if model == nil {
		t.Fatalf("expected model to be non-nil after loading example URDF")
	}
	if model.NumLinks() != 1 {
		t.Fatalf("expected 1 link from example URDF, got %d", model.NumLinks())
	}
}
