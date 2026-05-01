package loaders

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/WrenchRobotics/urdf-go/decoding"
	model_errors "github.com/WrenchRobotics/urdf-go/errors"
	urdfmodel "github.com/WrenchRobotics/urdf-go/urdf_model"
)

// FromURDFFile is a function which takes in a path to a URDF file and returns a
// pointer to a Model object.
//
// Expect an error if:
// - The file at the given path does not exist.
// - There was an issue reading the file at the given path.
// - There was an issue decoding the XML in the file at the given path.
// - No robot elements were found in the XML in the file at the given path.
// - There was an issue deriving the model from the robot element(s) found in the XML in the file at the given path.
//
// Example usage:
//
//	model, err := FromURDFFile("path/to/urdf/file.urdf")
//	if err != nil {
//	    // Handle the error
//	}
//	// Use the model
func FromURDFFile(path string) (*urdfmodel.Model, error) {
	// Setup

	// Check that the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("the file at path \"%v\" does not exist", path)
	}

	// Read the file's contents
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Decode
	var robotElt decoding.RobotElement
	err = xml.Unmarshal([]byte(content), &robotElt)
	if err != nil {
		return nil, fmt.Errorf("error decoding XML: %v", err)
	}

	// Ensure the root element is <robot>
	if robotElt.XMLName.Local != "robot" {
		return nil, model_errors.NoRobotsFoundInFileError{FilePath: path}
	}

	// Derive model
	model, err := urdfmodel.DeriveModelFrom(&robotElt)
	if err != nil {
		return nil, fmt.Errorf("there was an issue deriving the model: %v", err)
	}

	return model, nil

}
