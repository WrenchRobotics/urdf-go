package loaders

import (
	"encoding/xml"
	"fmt"

	"github.com/WrenchRobotics/urdf-go/decoding"
	model_errors "github.com/WrenchRobotics/urdf-go/errors"
	urdfmodel "github.com/WrenchRobotics/urdf-go/urdf_model"
)

// FromURDFContents is a function which takes in the contents of a URDF file and returns a
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
//	model, err := FromURDFContents([]byte("<robot>...</robot>"))
//	if err != nil {
//	    // Handle the error
//	}
//	// Use the model
func FromURDFContents(contents []byte) (*urdfmodel.Model, error) {
	// Decode the XML in the contents
	var robotElts []decoding.RobotElement
	err := xml.Unmarshal([]byte(contents), &robotElts)
	if err != nil {
		return nil, fmt.Errorf("error decoding XML: %v", err)
	}

	// Check that at least one robot element was found
	if len(robotElts) == 0 {
		return nil, model_errors.NoRobotsFoundInContentsError{}
	}

	// Derive model
	model, err := urdfmodel.DeriveModelFrom(&robotElts[0])
	if err != nil {
		return nil, fmt.Errorf("there was an issue deriving the model: %v", err)
	}

	return model, nil

}
