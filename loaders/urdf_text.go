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
// - There was an issue decoding the XML in the contents.
// - No robot elements were found in the XML in the contents.
// - There was an issue deriving the model from the robot element(s) found in the XML in the contents.
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
	var robotElt decoding.RobotElement
	err := xml.Unmarshal([]byte(contents), &robotElt)
	if err != nil {
		return nil, fmt.Errorf("error decoding XML: %v", err)
	}

	// Ensure the root element is <robot>
	if robotElt.XMLName.Local != "robot" {
		return nil, model_errors.NoRobotsFoundInContentsError{}
	}

	// Derive model
	model, err := urdfmodel.DeriveModelFrom(&robotElt)
	if err != nil {
		return nil, fmt.Errorf("there was an issue deriving the model: %v", err)
	}

	return model, nil

}
