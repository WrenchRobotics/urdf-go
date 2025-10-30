package urdfmodel

import (
	model_errors "github.com/WrenchRobotics/urdf-go/urdf_model/errors"
	"github.com/WrenchRobotics/urdf-go/urdf_model/joint"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link"
)

type Model struct {
	Name string
	// Unexported fields
	// (as a reminder, all fields with first character in lower case will not be exported outside of a Go package)
	rootLink  *link.Link
	links     map[string]*link.Link     // Complete list of links, organized by name
	joints    map[string]*joint.Joint   // Complete list of joints, organized by name
	materials map[string]*link.Material // Complete list of materials, organized by name
}

func (m Model) GetLink(linkName string) (*link.Link, error) {
	// Extract link pointer, if it exists
	out, ok := m.links[linkName]
	if !ok {
		return nil, model_errors.LinkDoesNotExistError{
			LinkName:  linkName,
			ModelName: m.Name,
		}
	}

	// Return successfully retrieved value
	return out, nil
}

func (m Model) GetJoint(jointName string) (*joint.Joint, error) {
	// Extract pointer to joint, if it exists
	out, ok := m.joints[jointName]
	if !ok {
		return nil, model_errors.JointDoesNotExistError{
			JointName: jointName,
			ModelName: m.Name,
		}
	}

	// Return successfully retrieved value
	return out, nil
}

func (m Model) GetMaterial(materialName string) (*link.Material, error) {
	// Extract pointer to material, if it exists
	out, ok := m.materials[materialName]
	if !ok {
		return nil, model_errors.MaterialDoesNotExistError{
			MaterialName: materialName,
			ModelName:    m.Name,
		}
	}

	// Return successfully retrieved material pointer
	return out, nil
}
