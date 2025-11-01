package loaders

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/WrenchRobotics/urdf-go/decoding"
	model_errors "github.com/WrenchRobotics/urdf-go/errors"
	urdfmodel "github.com/WrenchRobotics/urdf-go/urdf_model"
)

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
	var robotElts []decoding.RobotElement
	err = xml.Unmarshal([]byte(content), &robotElts)
	if err != nil {
		return nil, fmt.Errorf("error decoding XML: %v", err)
	}

	// Check that at least one robot element was found
	if len(robotElts) == 0 {
		return nil, model_errors.NoRobotsFoundError{FilePath: path}
	}

	// Derive model
	model, err := urdfmodel.DeriveModelFrom(&robotElts[0])
	if err != nil {
		return nil, fmt.Errorf("there was an issue deriving the model: %v", err)
	}

	return model, nil

}
